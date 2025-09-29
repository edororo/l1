package main

import (
	"fmt"
)

func main() {
	var n int64 // исходное число
	var i uint  // номер бита (от 0 до 63)
	var set int // 1 – установить бит в 1, 0 – сбросить бит

	// Ввод данных
	fmt.Println("Введите число: ")
	fmt.Scan(&n)

	fmt.Println("Введите номер бита: ")
	fmt.Scan(&i)

	fmt.Println("Введите действие (1 - установить бит в 1, 0 - установить бит в 0): ")
	fmt.Scan(&set)

	if set == 1 {
		n = n | (1 << i) // установка i-го бита в 1
	} else {
		n = n &^ (1 << i) // установка i-го бита в 0
	}

	fmt.Printf("Результат: %d (в двоичном виде: %064b)\n", n, uint64(n))
}
