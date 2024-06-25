package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Функция для быстрой сортировки массива
func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Выбираем опорный элемент
	pivot := arr[0]

	// Разделение массива на подмассивы по опорному элементу
	var low, high []int
	for _, v := range arr[1:] {
		if v <= pivot {
			low = append(low, v)
		} else {
			high = append(high, v)
		}
	}

	// Рекурсивная сортировка подмассивов
	quicksort(low)
	quicksort(high)

	// Слияние отсортированных подмассивов в исходный массив
	copy(arr, append(append(low, pivot), high...))
}

func main() {
	// Генерируем случайный массив чисел
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Исходный массив:", arr)

	// Вызываем быструю сортировку
	quicksort(arr)

	fmt.Println("Отсортированный массив:", arr)

	// Проверка с использованием встроенной сортировки
	arr2 := make([]int, len(arr))
	copy(arr2, arr)
	sort.Ints(arr2)
	fmt.Println("Отсортированный массив с использованием sort.Ints:", arr2)
}
