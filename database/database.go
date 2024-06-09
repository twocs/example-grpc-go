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
	user.ID = id // Set the ID to the value from the map.
	return &user, nil
}

// Search returns a list of users based on the query.
// No error is returned if no user is found.
func Search(query string) ([]*User, error) {
	var users []*User
	for id, user := range data {
		if user.Fname == query {
			user.ID = id // Set the ID to the value from the map.
			users = append(users, &user)
		}
	}
	if len(users) == 0 {
		return nil, nil
	}
	return users, nil
}
