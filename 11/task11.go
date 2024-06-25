package main

import (
	"fmt"
)

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 []int) []int {
	// Создаем map для первого множества для быстрого поиска
	set1Map := make(map[int]bool)
	for _, num := range set1 {
		set1Map[num] = true
	}

	// Создаем срез для хранения пересечения
	var result []int

	// Проверяем каждый элемент второго множества на наличие в первом множестве
	for _, num := range set2 {
		if set1Map[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	intersect := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", intersect)
}
