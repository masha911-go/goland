// **Задача 9.**
// Напишите программу, где одна горутина генерирует числа от 1 до 10 и отправляет их в канал. Другая горутина читает их
// и накапливает сумму. Используйте `sync.WaitGroup`, чтобы дождаться завершения обеих горутин перед выводом результата.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m int
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			ch <- i
		}
	}()
	go func(m *int) {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			*m += <-ch
		}
	}(&m)
	wg.Wait()
	fmt.Println(m)
}
