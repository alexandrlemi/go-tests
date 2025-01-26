package authserver

// TODO: JWT-токенизация
type TokenService interface {
	GenetareToken(userID string, attr string) (string, error)
	ValidateToken(token string) (bool, error)
}
