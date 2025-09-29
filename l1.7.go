package main

import (
	"fmt"
	"sync"
	"time"
)

var sm sync.Map

func main() {
	wg := sync.WaitGroup{}
	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go WWorker(&wg, w)
	}
	wg.Add(1)
	go RWorker(&wg)
	wg.Wait()
}

func WWorker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("%d", i)
		value := fmt.Sprintf("%d", i)
		sm.Store(key, value)
		time.Sleep(30 * time.Millisecond)
	}
}

func RWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(20 * time.Millisecond)
		fmt.Println("")
		sm.Range(func(key, value interface{}) bool {
			fmt.Printf("key: %s, value: %s\n", key, value)
			return true
		})
	}
}
