/*
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}

	Данный фрагмент кода может привести к проблемам с управлением памятью
	из-за использования большого объема памяти при создании огромной строки
	в функции createHugeString, переменная v выделяет память под строку
	размером 1 << 10 (1024 байт), что может быть избыточным и неэффективным
*/

package main

import (
	"fmt"
	"io"
	"strings"
)

var justString string

func someFunc() {
	// Используем буфер для чтения данных по частям
	const bufferSize = 1 << 10 // 1024 bytes
	buffer := make([]byte, bufferSize)

	// Пример чтения данных из createHugeString
	reader := strings.NewReader(createHugeString(bufferSize)) // Предположим, что функция createHugeString генерирует строку

	// Читаем первые 100 символов в justString
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка чтения данных:", err)
		return
	}

	justString = string(buffer[:n])
}

func main() {
	someFunc()
	fmt.Println("Первые 100 символов огромной строки:", justString)
}

func createHugeString(size int) string {
	return strings.Repeat("A", size)
}
