package main

import (
	"fmt"
	"time"
)

func worker(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Воркер заканчивает")
			return
		default:
			fmt.Println("Воркер работает...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan struct{})
	go worker(stopChan)

	time.Sleep(3 * time.Second)
	close(stopChan)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Программа завершена")
}
