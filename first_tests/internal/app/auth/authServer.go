package authserver

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

func (s *Authserver) Start(address string, port string) {

}

func NewServer() *Authserver {

	return &Authserver{}
}
