package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	wg := sync.WaitGroup{}
	ch := make(chan int)
	ch1 := make(chan int)
	wg.Add(3)
	go Write(&wg, nums, ch)
	go Read(&wg, ch, ch1)
	go Print(&wg, ch1)
	wg.Wait()
	fmt.Print("Контейнер завершён")

}

func Write(wg *sync.WaitGroup, nums [5]int, ch chan int) {
	defer wg.Done()
	for _, x := range nums {
		ch <- x
		time.Sleep(50 * time.Millisecond)
	}
	close(ch)
}

func Read(wg *sync.WaitGroup, ch chan int, ch1 chan int) {
	defer wg.Done()
	for x := range ch {
		ch1 <- x * 2
		time.Sleep(50 * time.Millisecond)

	}
	close(ch1)
}

func Print(wg *sync.WaitGroup, ch1 chan int) {
	defer wg.Done()
	for res := range ch1 {
		fmt.Println(res)
	}
}
