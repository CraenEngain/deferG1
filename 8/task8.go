package main

import (
	"fmt"
)

func main() {
	var num int64 = 19 // Исходное число
	bitIndex := 3      // Номер бита для установки

	// Установка i-го бита в 1
	num |= 1 << bitIndex

	fmt.Printf("Число после установки %d-го бита в 1: %d\n", bitIndex, num)

	// Установка i-го бита в 0
	num &^= 1 << bitIndex

	fmt.Printf("Число после установки %d-го бита в 0: %d\n", bitIndex, num)
}
