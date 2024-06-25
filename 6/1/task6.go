package main

import (
	"fmt"
	"time"
)

func worker(stopChan chan bool) {
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
	stopChan := make(chan bool)
	go worker(stopChan)

	time.Sleep(2 * time.Second)
	stopChan <- true
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Программа завершена")
}
