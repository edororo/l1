package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanln(&a, &b)
	x1, y1 := XOR(a, b)
	fmt.Println("XOR:", x1, y1)

	fmt.Scanln(&a, &b)
	x2, y2 := PlusMinus(a, b)
	fmt.Println("Сложение и вычитание:", x2, y2)
}

func XOR(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func PlusMinus(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
