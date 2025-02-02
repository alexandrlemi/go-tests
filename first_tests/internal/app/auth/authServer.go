package authserver

import "time"

type Authserver struct {
	UserRep     UserRepository
	TokenSrv    TokenService
	PsswdHasher PasswordHasher
	AlertSrv    AlertService
	LoggerSrv   LoggerService
	GRPCSrv     GRPCService
}

// Опишем используемые интерфейсы:
//
//	сохранения данных в DB
type Setter interface {
	Save(msg string)
}

// получения данных из DB
type Getter interface {
	Get(msg string)
}

// работы с базой данных
type UserRepository interface {
	Save(key string, msg string)
	Get(key string) string
}

// другой пакет
func NewRepoMock() UserRepository {
	return &repo{Stor: map[string]string{}}
}

type repo struct {
	Stor map[string]string
}

func (r *repo) Save(key string, msg string) {
	r.Stor[key] = msg
}
func (r *repo) Get(key string) string {
	return r.Stor[key]
}

// работы с токенами
type TokenService interface {
}

// шифрования паролей
type PasswordHasher interface {
}

// сервиса уведомлений
type AlertService interface {
}

// логгер
type LoggerService interface {
}

// GRPC
type GRPCService interface {
}

type trasport interface {
	refrash(func(refToken string) error) error
	//regist
	//login
	//refrash
}

func (s *Authserver) Start(tr trasport) {
	tr.refrash(s.refHandler())
}

func (s *Authserver) refHandler() func(refToken string) error {

	//
	time :=time.Now()

	return func(refToken string) error {
		println(time) 
		return nil
	}
}
func NewServer() *Authserver {

	return &Authserver{}
}
