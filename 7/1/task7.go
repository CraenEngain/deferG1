package main

import (
	"fmt"
	"sync"
)

func main() {
	var data sync.Map

	var wg sync.WaitGroup

	// Функция для записи данных в sync.Map
	write := func(key, value int) {
		defer wg.Done()
		data.Store(key, value)
	}

	// Запуск нескольких горутин для конкурентной записи данных в sync.Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write(i, i*i)
	}

	wg.Wait()

	// Вывод данных из sync.Map
	data.Range(func(key, value interface{}) bool {
		fmt.Printf("data[%d] = %d\n", key, value)
		return true
	})
}
