package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Используем big.Int для чисел
	a := big.NewInt(1123456789012345670)
	b := big.NewInt(1987654321098765430)
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Умножение a на b
	product := new(big.Int).Mul(a, b)
	fmt.Println("Умножение a * b =", product)

	// Деление a на b
	quotient := new(big.Int).Div(a, b)
	fmt.Println("Деление a / b =", quotient)

	// Сложение a и b
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сложение a + b =", sum)

	// Вычитание b из a
	difference := new(big.Int).Sub(a, b)
	fmt.Println("Вычитание a - b =", difference)
}
