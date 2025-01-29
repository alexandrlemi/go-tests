package main

import (
	authserver "first_test/internal/app/auth"
	"first_test/internal/app/db"
	"first_test/pkg/logger"
	"fmt"
	"os"
)

func main() {
	// Можно использовать конфиг

	log := logger.NewLogger(logger.LevelError, os.Stdout)

	userRepo := db.NewUserRepository()

	authServer := authserver.NewServer(userRepo, "secret")

	err := authServer.Register("mqkirill@mail.ru", "89044219992", "123qwe")
	if err != nil {
		log.Error("user", nil)
	}

	token, err := authServer.Login("mqkirill@mail.ru", "123qwe")
	if err != nil {
		log.Fatal("user not found", nil)
	}
	fmt.Println("TOKEN:", token)

	token, err = authServer.Login("89044219993", "123qwe")
	if err != nil {
		log.Fatal("user not found", nil)
	}
	fmt.Println("TOKEN:", token)

	//authServer.Start(address, port)

}
