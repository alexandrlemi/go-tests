package authserver

// TODO: Logger
type LoggerService interface {
	Info(msg string)
	Error(msg string)
}
