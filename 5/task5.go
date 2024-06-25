package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Время выполнения программы в секундах
	duration := 5 * time.Second

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Канал для передачи данных
	dataChan := make(chan int)

	// Горутина для отправки значений в канал
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				close(dataChan)
				return
			case dataChan <- i:
				i++
				time.Sleep(500 * time.Millisecond) // Имитация работы
			}
		}
	}()

	// Горутина для чтения значений из канала
	go func() {
		for val := range dataChan {
			fmt.Println("Получено:", val)
		}
	}()

	// Ожидание завершения контекста
	<-ctx.Done()
	fmt.Println("Программа завершена")
}
