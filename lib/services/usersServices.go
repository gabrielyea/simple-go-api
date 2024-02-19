package services

// MockUser represents a mock user structure
type MockUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUser returns a mock user
func GetUser() MockUser {
	return MockUser{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
}
