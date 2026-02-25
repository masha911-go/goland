// 9. Хранилище временных меток
// Создайте потокобезопасное хранилище последних временных меток событий, где:
// Запись происходит редко (обновление метки)
// Чтение происходит часто
package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	mu     sync.RWMutex
	events map[string]time.Time
}

func NewStore() *Store {
	return &Store{
		events: make(map[string]time.Time),
	}
}

func (s *Store) Set(event string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event] = time.Now()
	fmt.Printf("Событие '%s' обновлено\n", event)
}

func (s *Store) Get(event string) time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.events[event]
}

func main() {
	store := NewStore()

	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				t := store.Get("login")
				fmt.Printf("Читатель %d: %v\n", id, t.Format("15:04:05"))
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(2 * time.Second)
			store.Set("login")
		}
	}()

	time.Sleep(7 * time.Second)
}
