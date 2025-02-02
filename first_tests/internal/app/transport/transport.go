package transport

import "context"

type Transport interface {
	Register(ctx context.Context, email, phone, password string) error
	Login(ctx context.Context, identifier, password string) (string, error)
	RefreshToken(ctx context.Context, refreshToken string) error
}
