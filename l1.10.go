package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаём мапу: ключ — диапазон (десяток), значение — слайс температур
	groups := make(map[int][]float64)

	for _, t := range temps {
		// Находим диапазон, округляя вниз до ближайшего десятка
		key := int(math.Floor(t/10) * 10)
		groups[key] = append(groups[key], t)
	}

	// Выводим результат
	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, v)
	}
}