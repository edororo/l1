package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	done := make(chan struct{})
	ch := make(chan int)
	go Producer(&wg, ch, done)
	go Consumer(&wg, ch)
	go func() {
		<-time.After(7 * time.Second)
		fmt.Println("Завершение программы по таймеру")
		close(done)
	}()
	wg.Wait()
	fmt.Println("Программа завершена")
}

func Producer(wg *sync.WaitGroup, ch chan int, done chan struct{}) {
	defer wg.Done()
	defer close(ch)
	counter := 1
	for {
		select {
		case <-done:
			return
		case ch <- counter:
			ch <- counter
			fmt.Printf("Отправленное значение: %d\n", counter)
			counter++
			time.Sleep(21 * time.Millisecond)
		}
	}
}
func Consumer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for value := range ch {
		fmt.Printf("Полученное значение: %d\n", value)
		time.Sleep(21 * time.Millisecond)
	}
}
