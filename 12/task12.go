package main

import "fmt"

// Функция для создания множества строк
func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)

	// Добавляем каждую строку в множество
	for _, str := range strings {
		set[str] = true
	}

	return set
}

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество
	set := createSet(strings)

	// Выводим множество
	fmt.Println("Множество строк:")
	for key := range set {
		fmt.Println(key)
	}
}
