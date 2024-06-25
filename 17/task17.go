package main

import (
	"fmt"
)

// Функция для выполнения бинарного поиска в отсортированном массиве
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // Найден элемент
		} else if arr[mid] < target {
			low = mid + 1 // Искать в правой половине
		} else {
			high = mid - 1 // Искать в левой половине
		}
	}

	return -1 // Элемент не найден
}

func main() {
	// Пример отсортированного массива
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Элемент для поиска
	target := 24

	// Выполняем бинарный поиск
	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден в массиве по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
