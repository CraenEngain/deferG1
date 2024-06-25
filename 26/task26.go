package main

import (
	"fmt"
	"strings"
)

// Функция проверки уникальности символов в строке (регистронезависимая)
func isUnique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	// Инициализируем map для отслеживания встреченных символов
	charMap := make(map[rune]bool)

	// Проходим по каждому символу в строке
	for _, char := range str {
		// Если символ уже был встречен, возвращаем false
		if charMap[char] {
			return false
		}
		// Иначе отмечаем символ как встреченный
		charMap[char] = true
	}
	return true
}

func main() {
	// Тестовые примеры
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
	fmt.Println(isUnique("Hello"))
	fmt.Println(isUnique("uniqe"))
}
