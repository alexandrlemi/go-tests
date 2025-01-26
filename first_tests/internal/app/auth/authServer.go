package authserver

import (
	"context"
	"errors"
	db "first_test/internal/app/db"
	"golang.org/x/crypto/bcrypt"
)

type Authserver struct {
	UserRep   db.UserRepository
	TokenSrv  TokenService
	AlertSrv  AlertService
	LoggerSrv LoggerService
	GRPCSrv   GRPCService
	JWTKey    string
}

// TODO: Серверный слой
func (s *Authserver) Start(address string, port string) {

}

func NewServer(repo db.UserRepository, jwtKey string) *Authserver {

	return &Authserver{UserRep: repo, JWTKey: jwtKey}
}

// TODO: Авторизация
func (s *Authserver) Login(ctx context.Context, email, password string) (string, error) {
	hashedPassword, err := s.UserRep.Get(email)
	if err != nil {
		return "", errors.New("User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return "", errors.New("Incorrect password")
	}

	tokenString := ""
	return tokenString, nil
}

// TODO: Регистрация
func (s *Authserver) Register(ctx context.Context, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.UserRep.Save(email, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

// TODO: Работа с конфигом
// TODO: Слой QPA
