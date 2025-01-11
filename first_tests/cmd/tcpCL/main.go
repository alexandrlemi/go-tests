package main

import (
	"context"
	tcpcl "first_test/internal/app/tcpCL"
)

func main()  {
	//инициализируем лог
	//start клиента
	ctx:=context.Background()
	tc:=tcpcl.NewTcpCl()
	tc.Start(ctx,"localhost:8080")
}