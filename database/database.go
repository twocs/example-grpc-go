// Package database provides a mock database for testing purposes.
package database

import (
	"fmt"
)

// ErrUserNotFound is returned when the user ID is not found in the database.
var ErrUserNotFound = fmt.Errorf("user not found")

// User represents a user in the database.
type User struct {
	ID      int32  `json:"id"`
	Fname   string `json:"fname"` // First name
	City    string `json:"city"`
	Phone   string `json:"phone"`
	Height  string `json:"height"`
	Married bool   `json:"Married"`
}

// GetUserByID returns a user based on the user ID.
func GetUserByID(id int32) (*User, error) {
	user, ok := data[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return &user, nil
}
