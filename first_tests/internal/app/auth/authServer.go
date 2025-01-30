package authserver

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	db "first_test/internal/app/db"
	log "first_test/pkg/logger"
	"io"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrUserExists        = errors.New("user already exists")
)

type Authserver struct {
	UserRep   db.UserRepository //interface
	TokenSrv  TokenService
	AlertSrv  AlertService
	LoggerSrv log.Logger  //wrap over slog
	GRPCSrv   GRPCService
	JWTKey    string
	Pepper    string
}

// TODO: Серверный слой
func (s *Authserver) Start(address string, port string) {

}

func NewServer(repo db.UserRepository, jwtKey string) *Authserver {
	return &Authserver{UserRep: repo, JWTKey: jwtKey}
}

// TODO: Регистрация
//
//	Генерация соли
func generateSalt() (string, error) {
	saltBytes := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, saltBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(saltBytes), nil
}

func (s *Authserver) Register(email, phone, password string) error {
	_, _, errEmail := s.UserRep.Get(email)
	_, _, errPhone := s.UserRep.Get(phone)

	if errEmail == nil || errPhone == nil {
		return ErrUserExists
	}

	salt, err := generateSalt()
	if err != nil {
		return err
	}

	fullPassword := password + salt + s.Pepper
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fullPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.UserRep.Save(email, phone, string(hashedPassword), salt)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Авторизация
func (s *Authserver) Login(identifier, password string) (string, error) {
	// Нужно получить пароль и соль
	hashedPassword, salt, err := s.UserRep.Get(identifier)
	if err != nil {
		return "", ErrUserNotFound
	}
	fullPassword := password + salt + s.Pepper

	// Проверить хэш пароля
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(fullPassword)); err != nil {
		return "", ErrIncorrectPassword
	}

	tokenString := "JWT_Token"
	return tokenString, nil
}

// TODO: Работа с конфигом
// TODO: Слой QPA
