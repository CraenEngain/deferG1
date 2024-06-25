package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер заканчивает")
			return
		default:
			fmt.Println("Воркер работает...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Программа завершена")
}
