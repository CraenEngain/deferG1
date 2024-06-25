package main

import (
	"fmt"
	"sync"
	"time"
)

var stopFlag bool
var mu sync.Mutex

func worker() {
	for {
		mu.Lock()
		if stopFlag {
			mu.Unlock()
			fmt.Println("Воркер заканчивает")
			return
		}
		mu.Unlock()
		fmt.Println("Воркер работает...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	stopFlag = false
	go worker()

	time.Sleep(2 * time.Second)
	mu.Lock()
	stopFlag = true
	mu.Unlock()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Программа завершена")
}
