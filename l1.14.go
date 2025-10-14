package main

import "fmt"

func main() {
	a := make(chan bool)
	Type(a)
}

func Type(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип int")
	case bool:
		fmt.Println("Тип bool")
	case string:
		fmt.Println("Тип string")
	case chan string:
		fmt.Println("Тип chan string")
	case chan int:
		fmt.Println("Тип chan int")
	case chan bool:
		fmt.Println("Тип chan bool")
	default:
		fmt.Println("Неизвестный тип")
	}
}
