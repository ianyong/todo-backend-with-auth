package domainmodels

// User represents a person that uses the application.
type User struct {
	ID             int64
	Email          string
	HashedPassword string
}
