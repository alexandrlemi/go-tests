package tcpcl

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
)

//интерфейс для клиента
type TcpCL interface {
	Start(ctx context.Context, addres string) error
}


// конструктор возвращает интерфейс
func NewTcpCl() TcpCL {
	return &tcpCL{}
}

type tcpCL struct{
	conn net.Conn
}

func (tc *tcpCL)Start(ctx context.Context, addres string) error{
	
	conn,err:=net.Dial("tcp",addres)
	if err!=nil{
		log.Fatal("Не удалось подключиться к серверу")
		return err
	}
	defer conn.Close()

	tc.conn=conn

	for{
		reader:=bufio.NewReader(tc.conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by server")
			os.Exit(1)
		}
		fmt.Print(msg)
		fmt.Fprintf(tc.conn,"Привет от клиента \n")
	}

	return nil
}


func (tc *tcpCL)sendMSG(msg string)error{
	
	return nil
}
