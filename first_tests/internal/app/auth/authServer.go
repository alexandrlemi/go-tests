package authserver

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Authserver struct {
	userRep UserRepository
}

// Опишем используемые интерфейсы:

// работы с базой данных
type UserRepository interface {
	Save(key string, msg string) error
	Get(key string) error
}

type repo struct {
	Store map[string]string
}

func NewRepoMock() *repo {
	return &repo{Store: map[string]string{}}
}

func (r *repo) Save(key string, msg string) error {
	r.Store[key] = msg
	return nil
}
func (r *repo) Get(key string) error {
	return nil
}

type Transport interface {
	Refresh(func(refToken string) error) error
	Run() error
	Register(func(identifier, password string) error) error
}

func NewServer(repository UserRepository) *Authserver {

	return &Authserver{userRep: repository}
}

func (s *Authserver) Start(tr Transport) {
	tr.Refresh(s.refHandler())
	tr.Register(s.Register())
	tr.Run()
}

func (s *Authserver) refHandler() func(refToken string) error {

	//
	time := time.Now()

	return func(refToken string) error {
		fmt.Println(time)
		return nil
	}
}

func (s *Authserver) Register() func(identifier, password string) error {

	return func(identifier, password string) error {
		// TODO: Проверка, есть ли такой пользователь или нет
		err := s.userRep.Get(identifier)

		if err == nil {
			return fmt.Errorf("user not found")
		}

		// TODO: Генерация соли + перец из конфига

		fullPassword := password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fullPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		err = s.userRep.Save(identifier, string(hashedPassword))
		if err != nil {
			return err
		}

		return nil
	}
}
