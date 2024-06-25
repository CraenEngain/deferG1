package main

import (
	"fmt"
	"math"
)

// Определяем структуру Point
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p Point) DistanceTo(q Point) float64 {
	dx := q.x - p.x
	dy := q.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки с использованием конструктора
	p1 := NewPoint(1, 1)
	p2 := NewPoint(4, 5)

	// Вычисляем расстояние между точками
	distance := p1.DistanceTo(p2)

	// Выводим результат
	fmt.Printf("Расстояние между точками p1(%g, %g) и p2(%g, %g) равно %g\n", p1.x, p1.y, p2.x, p2.y, distance)
}
