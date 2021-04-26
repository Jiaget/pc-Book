package service

import "golang.org/x/crypto/bcrypt"

// User contains user's information
type User struct {
	UserName       string
	HashedPassword string
	Role           string
}

// NewUser returns a new user
func NewUser(username string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		UserName:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
	}
	return user, nil
}

// IsCorrectPassword checks if the provided password is correct
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

// Clone returns a cloned user
func (user *User) Clone() *User {
	return &User{
		UserName:       user.UserName,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
	}
}
