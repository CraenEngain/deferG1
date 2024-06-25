package main

import (
	"fmt"
	"os"
)

func main() {
	// Исходный слайс
	slice := []int{1, 2, 5, 4, 22}

	// Индекс элемента, который нужно удалить
	var index int
	fmt.Sscanf(os.Args[1], "%d", &index)

	// Удаление элемента из слайса
	if index < len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
	}

	// Вывод слайса после удаления элемента
	fmt.Println(slice)
}
