// 8. Graceful shutdown флаг
// Реализуйте механизм graceful shutdown, где:
// Одна горутина может установить флаг завершения
// Множество горутин могут проверять этот флаг
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type ShutdownFlag struct {
	flag int32
}

func (sf *ShutdownFlag) Shutdown() {
	atomic.StoreInt32(&sf.flag, 1)
}

func (sf *ShutdownFlag) IsShutdown() bool {
	return atomic.LoadInt32(&sf.flag) == 1
}

func main() {
	var shutdown ShutdownFlag
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for !shutdown.IsShutdown() {
				fmt.Printf("Горутина %d работает\n", id)
				time.Sleep(500 * time.Millisecond)
			}
			fmt.Printf("Горутина %d завершилась\n", id)
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Завершаем работу")
	shutdown.Shutdown()

	wg.Wait()
	fmt.Println("Все горутины завершены, программа закрывается")
}
