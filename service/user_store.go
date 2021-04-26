package service

import "sync"

// UserStore is an interface to store users
type UserStore interface {
	// Save saves the user into the store
	Save(user *User) error
	// Find finds the user from the store by username
	Find(username string) (*User, error)
}

// InMemoryUserStore stores the users in memory
type InMemoryUserStore struct {
	mutex sync.Mutex
	users map[string]*User
}

// NewInMemoryUserStore returns a new InMemoryUserStore
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*User),
	}
}

func (store *InMemoryUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.UserName] != nil {
		return ErrAlreadyExists
	}

	// we cannot save the user directly, for the user here is a point type
	store.users[user.UserName] = user.Clone()
	return nil
}

func (store *InMemoryUserStore) Find(username string) (*User, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	user := store.users[username]
	if user != nil {
		return user.Clone(), nil
	}
	return nil, nil
}
