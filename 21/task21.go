package main

import "fmt"

// MetricCalculator представляет калькулятор для метрических измерений
type MetricCalculator struct{}

// CalculateDistance возвращает расстояние в метрах
func (m *MetricCalculator) CalculateDistance() float64 {
	return 1000 // Предположим, что это расстояние в метрах
}

// EnglishCalculator интерфейс для работы с английскими измерениями
type EnglishCalculator interface {
	CalculateDistanceInFeet() float64
}

// MetricToEnglishAdapter адаптер для конвертации метрических измерений в английские
type MetricToEnglishAdapter struct {
	calculator *MetricCalculator
}

// NewMetricToEnglishAdapter создает новый адаптер для конвертации из метрической системы в английскую
func NewMetricToEnglishAdapter(calculator *MetricCalculator) *MetricToEnglishAdapter {
	return &MetricToEnglishAdapter{calculator: calculator}
}

// CalculateDistanceInFeet конвертирует расстояние из метров в футы
func (a *MetricToEnglishAdapter) CalculateDistanceInFeet() float64 {
	// Получаем расстояние в метрах с помощью метрического калькулятора
	distanceInMeters := a.calculator.CalculateDistance()

	// Конвертируем метры в футы (1 метр = 3.28084 фута)
	const metersToFeet = 3.28084
	distanceInFeet := distanceInMeters * metersToFeet

	return distanceInFeet
}

func main() {
	// Создаем экземпляр метрического калькулятора
	metricCalculator := &MetricCalculator{}

	// Создаем адаптер для конвертации метрических измерений в английские
	adapter := NewMetricToEnglishAdapter(metricCalculator)

	// Используем адаптер для расчета расстояния в футах
	distanceInFeet := adapter.CalculateDistanceInFeet()

	// Выводим результат
	fmt.Printf("Расстояние в футах: %.2f\n", distanceInFeet)
}
