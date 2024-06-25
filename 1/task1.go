package main

import "fmt"

// Определим структуру Human
type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Printf("Меня зовут %s. Мне %d лет.\n", h.Name, h.Age)
}

// Определим структуру Action
type Action struct {
	Human
	Sex string
}

// Метод для структуры Action
func (a Action) Recognition() {
	fmt.Printf("%s - %s\n", a.Name, a.Sex)
}

func main() {
	h := Human{Name: "Мария", Age: 36}
	h.Speak()

	a := Action{
		Human: Human{Name: "Антон", Age: 28},
		Sex:   "Мужчина",
	}
	a.Speak()       // Вызовем метод Speak из встроенной структуры Human
	a.Recognition() // Вызовем метод Recognition структуры Action
}
