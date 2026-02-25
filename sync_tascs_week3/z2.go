// 2. Атомарный счетчик
// Создайте потокобезопасный счетчик, который может инкрементироваться и
// возвращать текущее значение, используя только атомарные операции (atomic).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int32
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
