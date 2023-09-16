package model

import (
	"crypto/sha256"
	"fmt"
)

// User is a user of the application
type User struct {
	Id       string
	Name     string
	Email    string
	Password string
	Token    string
	IsAdmin  bool
}

// HashPassword obfuscates a user's password
// using a one-way sha256 hash
func (u *User) HashPassword() error {
	hash := sha256.New()
	_, err := hash.Write([]byte(u.Id + ":" + u.Email + ":" + u.Password))
	if err != nil {
		return err
	}
	u.Password = fmt.Sprintf("%x", hash.Sum(nil))
	return nil
}
