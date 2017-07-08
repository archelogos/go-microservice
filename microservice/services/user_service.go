package services

import "log"

// User ...
type User struct {
	Name string
}

// GetUser ...
func GetUser(name string) User {
	log.Println("[SERVICE] Getting user")
	return User{
		Name: name,
	}
}
