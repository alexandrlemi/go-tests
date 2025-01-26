package authserver

// TODO: Двух-факторная аутентификация
type AlertService interface {
	SendCode(userID string) error
	ValidateCode(userID string, code string) (bool, error)
}
