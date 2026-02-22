// **Задача 2.**
// Реализуйте структуру `SafeCounter` с методами:
// - `Inc()` — увеличивает внутренний счётчик;
// - `Value() int` — возвращает текущее значение.
// Используйте `sync.Mutex`. Протестируйте её с 50 параллельными вызовами `Inc()`.
package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (sc *SafeCounter) Inc() {
	sc.mu.Lock()
	sc.count++
	sc.mu.Unlock()
}

func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

func main() {
	counter := new(SafeCounter)
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			counter.Inc()
		}(i)
	}

	wg.Wait()
	fmt.Println("Result:", counter.Value())
}
