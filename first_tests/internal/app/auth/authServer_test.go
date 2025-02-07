package authserver_test

///

import (
	auth "first_test/internal/app/auth"
	transport "first_test/internal/app/transportRestTest"
	"testing"
)

// Тесты

func TestTransportRefrash(t *testing.T) {
	tr := transport.NewTransport(":8080")
	authsrv := auth.Authserver{}
	authsrv.Start(tr)
	
	//TODO сделать запрос прям от сюда

}

func TestAuthserver_Register(t *testing.T) {
	tr := transport.NewTransport(":8080")
	mockRepo := auth.NewRepoMock()
	authSrv := auth.NewServer(mockRepo)
	authSrv.Start(tr)

}
