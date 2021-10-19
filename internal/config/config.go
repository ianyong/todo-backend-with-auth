package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	ServerPort  int
	SecretKey   string
	DBHost      string
	DBPort      int
	DBName      string
	DBUser      string
	DBPassword  string
	DBSSLMode   string
}

const (
	envDevelopment        = "development"
	envTest               = "test"
	envProduction         = "production"
	configLoadErrTemplate = "unable to load env vars: %w"
)

var (
	environments        = []string{envDevelopment, envTest, envProduction}
	developmentEnvFiles = []string{".env.development.local", ".env.development"}
	testEnvFiles        = []string{".env.test.local", ".env.test"}
	productionEnvFiles  = []string{".env"}
	projectRootPath     = getProjectRootPath()
)

// Load loads the application configuration from environment variables.
func Load() (*Config, error) {
	var cfg Config

	err := cfg.loadEnvironment()
	if err != nil {
		return nil, fmt.Errorf(configLoadErrTemplate, err)
	}

	err = cfg.loadConfig()
	if err != nil {
		return nil, fmt.Errorf(configLoadErrTemplate, err)
	}

	return &cfg, nil
}

// LoadTest loads the application configuration for running tests.
func LoadTest() (*Config, error) {
	var cfg Config
	cfg.Environment = envTest

	err := cfg.loadConfig()
	if err != nil {
		return nil, fmt.Errorf(configLoadErrTemplate, err)
	}

	return &cfg, nil
}

// loadEnvironment loads the GO_ENV environment variable from the shell environment.
// If GO_ENV is not set, it defaults to the 'development' environment.
func (cfg *Config) loadEnvironment() error {
	env := os.Getenv("GO_ENV")
	if env == "" {
		cfg.Environment = envDevelopment
		return nil
	}
	for _, environment := range environments {
		if env == environment {
			cfg.Environment = env
			return nil
		}
	}
	return fmt.Errorf("unrecognised environment: %s", env)
}

// loadConfig loads all fields of Config except Environment.
// The sequence of files that are loaded depends on the value of Environment.
func (cfg *Config) loadConfig() error {
	var configFileNames []string
	switch cfg.Environment {
	case envDevelopment:
		configFileNames = developmentEnvFiles
	case envTest:
		configFileNames = testEnvFiles
	case envProduction:
		configFileNames = productionEnvFiles
	}
	err := loadEnvFromFiles(configFileNames)
	if err != nil {
		log.Println("Loading configuration from shell environment...")
	}

	cfg.ServerPort, err = parsePort(os.Getenv("SERVER_PORT"))
	if err != nil {
		return fmt.Errorf("unable to parse server port: %w", err)
	}
	cfg.SecretKey = os.Getenv("SECRET_KEY")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort, err = parsePort(os.Getenv("DB_PORT"))
	if err != nil {
		return fmt.Errorf("unable to parse db port: %w", err)
	}
	cfg.DBName = os.Getenv("DB_NAME")
	if cfg.DBName == "" {
		return errors.New("unable to parse db name: cannot be empty")
	}
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBSSLMode = os.Getenv("DB_SSLMODE")

	return nil
}

// getProjectRootPath returns the root path of the project dynamically. This is in contrast to hard-coded paths
// which break upon building the server into an executable.
func getProjectRootPath() string {
	// Get full path of current file from runtime.
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("unable to retrieve root path of server")
	}

	// Traverse two parent directories up. Note that this is sensitive to the level of nesting
	// of this file within the project structure.
	return filepath.Join(filepath.Dir(b), "../../")
}

// loadEnvFromFile reads in environment variables declared within a file.
func loadEnvFromFile(configFileName string) error {
	fullFilePath := fmt.Sprintf("%s/%s", projectRootPath, configFileName)
	err := godotenv.Load(fullFilePath)
	if err != nil {
		return fmt.Errorf("error while loading %s: %w", configFileName, err)
	}
	return nil
}

// loadEnvFromFiles attempts to read in environment variables declared within a list of files sequentially
// until a file is successfully read.
func loadEnvFromFiles(configFileNames []string) error {
	for _, filename := range configFileNames {
		log.Printf("Loading configuration from '%s'...\n", filename)
		err := loadEnvFromFile(filename)
		// If a config file managed to load, do not load the other files.
		if err == nil {
			return nil
		}
		log.Printf("'%s' not found. Loading next config file...\n", filename)
	}

	configFileNamesList := strings.Join(configFileNames, ", ")
	return fmt.Errorf("no config file found from any of %s", configFileNamesList)
}

// parsePort converts a string into an integer if it is within the acceptable range for port numbers (0 to 65535).
func parsePort(value string) (int, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("invalid port: not an integer")
	}
	if result < 0 || result > 65535 {
		return 0, errors.New("invalid port: port number should be a 16-bit unsigned integer")
	}
	return result, nil
}
