package main

import (
	"fmt"
	"reflect"
)

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Переменная является типом int")
	case string:
		fmt.Println("Переменная является типом string")
	case bool:
		fmt.Println("Переменная является типом bool")
	case chan int:
		fmt.Println("Переменная является типом chan int")
	default:
		fmt.Println("Не удалось определить тип переменной")
	}
}

func main() {
	// Примеры переменных различных типов
	var v1 interface{} = 42
	var v2 interface{} = "hello"
	var v3 interface{} = true
	var v4 interface{} = make(chan int)

	// Определение типов переменных
	checkType(v1)
	checkType(v2)
	checkType(v3)
	checkType(v4)

	// Использование рефлексии в случае использование динамических типов
	fmt.Printf("Тип переменной v1: %s\n", reflect.TypeOf(v1))
}
