package main

import (
	"fmt"
)

// Функция для переворачивания строки с учетом Unicode символов
func reverseString(input string) string {
	// Преобразуем строку в массив рун (Unicode символов)
	runes := []rune(input)
	// Инвертируем порядок символов в массиве рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем массив рун обратно в строку
	return string(runes)
}

func main() {
	str := "главрыба"
	fmt.Println("Исходная строка:", str)

	// Переворачиваем строку
	reversed := reverseString(str)
	fmt.Println("Перевернутая строка:", reversed)
}
