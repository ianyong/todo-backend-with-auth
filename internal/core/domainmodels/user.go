package domainmodels

// User represents a person that uses the application.
type User struct {
	Email          string
	HashedPassword string
}
