package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	fmt.Println("Первый способ - счётчик")
	ExitCondition(wg)
	fmt.Println("\nВторой способ - сигнал уведомления")
	DoneChannel()
	fmt.Println("\nТретий способ - context")
	Context()
	fmt.Println("\nЧетвёртый способ - runtime.Goexit()")
	GoExit()
	fmt.Println("\nПятый способ - time.After")
	TimeAfter()
	fmt.Println("\nШестой способ - Ctrl+C")
	Sigint()
	fmt.Println("Седьмой способ - закрытие канала")
	channelClose()
}

func ExitCondition(wg *sync.WaitGroup) {
	counter := 1
	defer wg.Done()
	for {
		fmt.Println("Первый способ: ", counter)
		counter++
		time.Sleep(500 * time.Millisecond)
		if counter == 6 {
			fmt.Println("Способ первый по счётчику - завершён")
			break
		}
	}
}

func DoneChannel() {
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		counter := 1
		for {
			select {
			case <-done:
				fmt.Println("Получен сигнал остановки - второй способ завершён")
				return
			default:
				fmt.Println("Второй способ: ", counter)
				counter++
				time.Sleep(650 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second) // через 3 сек. подаём сигнал
	close(done)
	wg.Wait() // ждём завершения горутины
}

func Context() {
	// Контекст с ручной отменой (WithCancel)
	fmt.Println("WithCancel")
	ctxCancel, cancel := context.WithCancel(context.Background())
	go worker("CancelWorker", ctxCancel)
	time.Sleep(2 * time.Second)
	fmt.Println("Отправляем cancel()")
	cancel()
	time.Sleep(1 * time.Second)

	// Контекст с таймаутом (WithTimeout)
	fmt.Println("\nДемонстрация WithTimeout")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelTimeout()
	go worker("TimeoutWorker", ctxTimeout)
	time.Sleep(4 * time.Second) // ждём автомат.таймаута

	// Контекст с дедлайном (WithDeadline)
	fmt.Println("\nДемонстрация WithDeadline ")
	deadline := time.Now().Add(2 * time.Second)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	defer cancelDeadline()
	go worker("DeadlineWorker", ctxDeadline)
	time.Sleep(3 * time.Second) // ждём дедлайна
}

func worker(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s Завершение работы: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s Работаю...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func GoExit() {
	go func() {
		defer fmt.Println("Просто пример, что дефер выполнится")

		fmt.Println("Четвёртый способ начался")
		time.Sleep(1 * time.Second)
		fmt.Println("Заканчивается")
		runtime.Goexit()
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Четвёртый способ через runtime - завершён")
}

func TimeAfter() {
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for i := 1; ; i++ {
			select {
			case <-time.After(4 * time.Second):
				fmt.Println("TimeAfter: время вышло, завершаем")
				return
			case <-ticker.C:
				fmt.Println("Пятый способ:", i)
			}
		}
	}()
	time.Sleep(2 * time.Second)
}

func Sigint() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt) // подписка на Ctrl+C

	done := make(chan struct{})
	go func() {
		counter := 1
		for {
			select {
			case <-done:
				fmt.Println("SigintWorker: получен сигнал завершения")
				return
			default:
				fmt.Println("SigintWorker:", counter)
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	fmt.Println("Нажмите Ctrl+C для остановки")
	<-sig
	close(done)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Шестой способ завершён")
}

func channelClose() {
	ch := make(chan int)

	go func() {
		for val := range ch {
			fmt.Println("Worker получил:", val)
			time.Sleep(400 * time.Millisecond)
		}
		fmt.Println("Worker завершил работу (канал закрыт)")
	}()

	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(300 * time.Millisecond)
	}

	close(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("Седьмой способ завершён")
}
