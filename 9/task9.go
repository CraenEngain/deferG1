package main

import (
	"fmt"
	"sync"
)

func main() {
	// Массив чисел для конвейера
	numbers := []int{1, 2, 3, 4, 5}

	// Создаем каналы для конвейера
	input := make(chan int)
	output := make(chan int)

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Горутина для записи чисел в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	// Горутина для чтения чисел из первого канала, умножения на 2 и записи результата во второй канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range input {
			output <- num * 2
		}
		close(output)
	}()

	// Горутина для вывода данных из второго канала в stdout
	go func() {
		for result := range output {
			fmt.Println(result)
		}
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
}
