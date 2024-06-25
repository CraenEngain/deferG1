package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Функция воркера, который читает данные из канала и выводит их в stdout
func worker(ctx context.Context, id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d заканчивает\n", id)
			return
		case data := <-dataChan:
			fmt.Printf("Воркер %d получил данные: %d\n", id, data)
		}
	}
}

func main() {
	dataChan := make(chan int)
	var wg sync.WaitGroup

	// Создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Чтение количества воркеров из аргументов командной строки
	var numWorkers int
	fmt.Sscanf(os.Args[1], "%d", &numWorkers)

	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataChan, &wg)
	}

	// Запускаем горутину для генерации данных
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dataChan)
				return
			default:
				dataChan <- rand.Intn(100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Перехватываем сигнал Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнал
	<-c
	fmt.Println("\nReceived shutdown signal, stopping workers...")
	cancel()

	// Ждем завершения всех воркеров
	wg.Wait()
	fmt.Println("All workers stopped. Exiting.")
}
