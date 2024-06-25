package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	data := make(map[int]int)

	var wg sync.WaitGroup

	// Функция для записи данных в map
	write := func(key, value int) {
		defer wg.Done()
		mu.Lock()
		data[key] = value
		mu.Unlock()
	}

	// Запуск нескольких горутин для конкурентной записи данных в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write(i, i*i)
	}

	wg.Wait()

	// Вывод данных из map
	mu.Lock()
	for key, value := range data {
		fmt.Printf("data[%d] = %d\n", key, value)
	}
	mu.Unlock()
}
