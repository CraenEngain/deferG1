package main

import (
	"fmt"
	"sync"
)

// Функция для вычисления квадрата числа и отправки результата в канал
func square(num int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done() // Сообщаем, что горутина завершилась
	square := num * num
	resultChan <- square // Отправляем результат в канал
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	resultChan := make(chan int, len(numbers)) // Канал для результатов

	// Запускаем горутины для вычисления квадратов чисел
	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, resultChan)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Читаем и выводим результаты из канала
	for result := range resultChan {
		fmt.Println(result)
	}
}
