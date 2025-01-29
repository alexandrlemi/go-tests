package db

import (
	"errors"
	"sync"
)

type User struct {
	Email          string
	Phone          string
	hashedPassword string
	Salt           string
}

var DATABASE = struct {
	sync.Mutex
	usersByEmail map[string]User
	usersByPhone map[string]User
}{
	usersByEmail: make(map[string]User),
	usersByPhone: make(map[string]User),
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserExists = errors.New("user already exists")

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r *UserRepository) Save(email, phone, hashedPassword, salt string) error {
	DATABASE.Lock()
	defer DATABASE.Unlock()

	if _, exist := DATABASE.usersByEmail[email]; exist {
		return ErrUserExists
	}
	if _, exist := DATABASE.usersByPhone[phone]; exist {
		return ErrUserExists
	}

	user := User{
		Email:          email,
		Phone:          phone,
		hashedPassword: hashedPassword,
		Salt:           salt,
	}
	DATABASE.usersByEmail[email] = user
	DATABASE.usersByPhone[phone] = user
	return nil
}

func (r *UserRepository) Get(identifier string) (string, string, error) {
	DATABASE.Lock()
	defer DATABASE.Unlock()

	user, exist := DATABASE.usersByEmail[identifier]
	if !exist {
		user, exist = DATABASE.usersByPhone[identifier]
		if !exist {
			return "", "", ErrUserNotFound
		}
	}
	return user.hashedPassword, user.Salt, nil
}
