package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		defer wg.Done()
		counter := 1

		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- counter
				fmt.Printf("Отправленное значение: %d\n", counter)
				counter++
				time.Sleep(20 * time.Millisecond)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Printf("Полученное значение: %d\n", value)
			time.Sleep(20 * time.Millisecond)
		}
	}()
	wg.Wait()
	fmt.Println("Программа завершена по Ctrl + C")
}
