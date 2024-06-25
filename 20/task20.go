package main

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Инвертируем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова в строку с пробелами между ними
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println("Исходная строка:", str)

	// Переворачиваем слова в строке
	reversed := reverseWords(str)
	fmt.Println("Перевернутая строка:", reversed)
}
