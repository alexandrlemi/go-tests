package tcpsrv

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"net"
)

// интерфейс для клиента
type TcpSrv interface {
	Start(ctx context.Context, addres string) error
}
git
// конструктор возвращает интерфейс
func NewTcpSrv(logger *slog.Logger) TcpSrv {

	return &tcpSrv{logger: logger}
}

type tcpSrv1 struct {
	listener net.Listener
	logger   *slog.Logger
}

func (ts *tcpSrv1) Start(ctx context.Context, addres string) error {
	return nil
}

type tcpSrv struct {
	listener net.Listener
	logger   *slog.Logger
}

func (ts *tcpSrv) Start(ctx context.Context, addres string) error {

	listener, err := net.Listen("tcp", addres)
	if err != nil {
		ts.logger.Error("Не удалось подключиться к серверу")
		return err
	}
	defer listener.Close()
	ts.listener = listener
	ts.logger.Info("Сервер запущен ", addres)

	go ts.listenerLoop()

	select {
	case <-ctx.Done(): // Если контекст отменён
		ts.logger.Error("Контекст отменён:", ctx.Err())
		return nil
	}
}

func (ts *tcpSrv) listenerLoop() {
	for {
		conn, err := ts.listener.Accept()
		if err != nil {
			ts.logger.Error("Ошибка подключения:", err)
			continue
		}
		go ts.handleConnection(conn)
	}
}

func (ts *tcpSrv) handleConnection(conn net.Conn) {
	for {
		fmt.Fprintf(conn, "Привет\n")
		reader := bufio.NewReader(conn)
		tmp, err := reader.ReadString('\n')
		if err != nil {
			ts.logger.Error("Ошибка при чтении: ", err)
			return
		}
		ts.logger.Info(string(tmp))
	}
}
