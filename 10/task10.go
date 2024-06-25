package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем map для группировки температур по диапазонам
	groupedTemperatures := make(map[int][]float64)

	// Обходим все температуры и добавляем их в соответствующую группу
	for _, temp := range temperatures {
		key := int(math.Round(temp)) / 10 * 10
		groupedTemperatures[key] = append(groupedTemperatures[key], temp)
	}

	// Выводим результаты
	for key, temps := range groupedTemperatures {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
