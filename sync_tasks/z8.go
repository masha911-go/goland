// **Задача 8.**
// Создайте структуру `Stats` с полями `requests` и `errors` (оба `int`). Добавьте методы `RecordRequest()` и
// `RecordError()`, защищённые одним мьютексом. Запустите 20 горутин, имитирующих запросы (90% успех, 10% ошибка).
// Проверьте итоговую статистику.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Stats struct {
	mu       sync.Mutex
	requests int
	errors   int
}

func (s *Stats) RecordRequest() {
	s.mu.Lock()
	s.requests++
	s.mu.Unlock()

}

func (s *Stats) RecordError() {
	s.mu.Lock()
	s.errors++
	s.mu.Unlock()

}

func main() {
	rand.Seed(time.Now().UnixNano())
	stats := &Stats{}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				stats.RecordRequest()
				if rand.Float64() < 0.1 {
					stats.RecordError()
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Всего запросов: %d\n", stats.requests)
	fmt.Printf("Всего ошибок: %d\n", stats.errors)
}
