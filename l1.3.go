package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	// проверяем кол-во параметров при запуске

	if len(os.Args) < 2 {
		return
	}
	// читаем
	workers, err := strconv.Atoi(os.Args[1])
	if err != nil || workers <= 0 {
		fmt.Println("Неверный параметр количества воркеров")
		return
	}
	for i := 0; i <= workers; i++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for k := range ch {
				fmt.Printf("Воркер %d, данные %v\n", worker, k)
				time.Sleep(400 * time.Millisecond)
			}
		}(i)
	}
	counter := 1
	for {
		ch <- counter
		fmt.Printf("Данные записаны %d\n", counter)
		counter++
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}
