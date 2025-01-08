package main

import (
	"context"
	"fmt"
	"time"
)

// Функция, которая выполняет какую-то работу и сигнализирует о завершении через контекст
func worker(ctx context.Context,cancel context.CancelFunc, done chan<- struct{}) {
	defer close(done) // Закрываем канал при завершении работы
	//test Nikita2 
	fmt.Println("Горутина начала работу")
	
	time.Sleep(1 * time.Second) // Имитация работы
	cancel()
	time.Sleep(1 * time.Second) // Имитация работы

	fmt.Println("Горутина завершила работу")

	// Сигнализируем о завершении
	done <- struct{}{}
}

func main() {
	// Создаём контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Отменяем контекст при завершении main

	// Канал для сигнала о завершении работы горутины
	done := make(chan struct{})

	// Запускаем горутину
	go worker(ctx, cancel,done)

	// Ожидаем завершения горутины
	fmt.Println("Ожидаем завершения горутины...")
	select {
	case <-done: // Ждём сигнала о завершении
		fmt.Println("Горутина завершила работу, main продолжает выполнение")
	case <-ctx.Done(): // Если контекст отменён
		fmt.Println("Контекст отменён:", ctx.Err())
	}

	fmt.Println("Программа завершена.")
}