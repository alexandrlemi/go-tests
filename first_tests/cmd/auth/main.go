package main

import (
	authserver "first_test/internal/app/auth"
	"log/slog"
)

func main() {
	//start
	
	ss:=authserver.SeverStruct{}
	authServer:=authserver.NewServer(slog.New(slog.Default().Handler()),&ss)
	authServer.Start()
	
}