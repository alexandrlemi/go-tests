package main

import (
	authserver "first_test/internal/app/auth"
)

func main() {
	//start
	// Можно использовать конфиг
	address := "localhost"
	port := "50051"
	authServer := authserver.NewServer()
	authServer.Start(address, port)

}
