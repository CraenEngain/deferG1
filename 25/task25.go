package main

import (
	"fmt"
	"os"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Начало программы")
	var delay int
	fmt.Sscanf(os.Args[1], "%d", &delay)
	// Вызов функции sleep на 3 секунды
	sleep(time.Duration(delay) * 1e9)
	fmt.Println("Программа завершена")
}
