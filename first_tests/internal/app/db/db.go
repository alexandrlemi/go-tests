package db

import "errors"

type UserRepository interface {
	Save(email string, password string) error
	Get(email string) (string, error)
}

type InMemoryRepository struct {
	data map[string]string
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: make(map[string]string)}
}

func (r *InMemoryRepository) Save(email string, password string) error {
	if _, exists := r.data[email]; exists {
		return errors.New("email already exists")
	}
	r.data[email] = password
	return nil
}

func (r *InMemoryRepository) Get(email string) (string, error) {
	password, exists := r.data[email]
	if !exists {
		return "", errors.New("email not found")
	}
	return password, nil
}
