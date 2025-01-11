package main

import (
	"context"
	tcpsrv "first_test/internal/app/tcpSrv"
	"log/slog"
	"os"
)

func main() {
	//инициализируем лог
	handler:=slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{Level: slog.LevelDebug})
	log:=slog.New(handler)
	//start клиента
	
	ctx,cancel := context.WithCancel(context.Background())  
	defer cancel()

	ts111 := tcpsrv.NewTcpSrv(log)
	ts111.Start(ctx, "localhost:8080")

}