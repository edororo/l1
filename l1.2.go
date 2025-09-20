package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Квадрат %v: %d\n", n, number*number)
		}(number)
	}
	wg.Wait()
}
