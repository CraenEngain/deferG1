package main

import "fmt"

func main() {
	a, b := 5, 10

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
