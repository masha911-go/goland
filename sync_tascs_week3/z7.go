// 7. Кэш статистики с редкими обновлениями
// Создайте кэш статистических данных, где:
// Данные обновляются раз в минуту (одна горутина)
// Данные читаются сотни раз в секунду (много горутин)
package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

func main() {
	cache := &Cache{
		data: map[string]int{
			"views": 1000,
			"likes": 500,
		},
	}

	go func() {
		for range time.Tick(time.Minute) {
			cache.mu.Lock()
			cache.data["views"] += 100
			cache.mu.Unlock()
			fmt.Println("Updated")
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for range time.Tick(10 * time.Millisecond) {
				cache.mu.RLock()
				_ = cache.data["views"]
				cache.mu.RUnlock()
			}
		}()
	}

	time.Sleep(90 * time.Second)
}
