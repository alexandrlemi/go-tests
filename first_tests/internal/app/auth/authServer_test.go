package authserver_test

///

import (
	auth "first_test/internal/app/auth"
	transport "first_test/internal/app/transportRestTest"
	"testing"
)

// Тесты




func TestTransportRefrash(t *testing.T) {
	tr:= transport.NewTransport(":8080")
	tr2:= transport.NewTransport(":8081")

	authsrv:=auth.Authserver{}
	go	authsrv.Start(tr)
	authsrv.Start(tr2)
	//TODO сделать запрос прям от сюда
	
	
}