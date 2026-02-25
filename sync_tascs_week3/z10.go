// 10. Rate limiter на atomic
// Реализуйте простой ограничитель запросов (rate limiter), который:
// Считает количество запросов в текущем интервале
// Сбрасывает счетчик по таймеру
// Должен работать максимально эффективно при проверках (atomic).

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	counter  int64
	limit    int64
	interval time.Duration
	resetCh  chan struct{}
}

func NewRateLimiter(limit int64, interval time.Duration) *RateLimiter {
	rl := &RateLimiter{
		limit:    limit,
		interval: interval,
		resetCh:  make(chan struct{}),
	}

	go rl.resetLoop()
	return rl
}

func (rl *RateLimiter) Allow() bool {
	current := atomic.AddInt64(&rl.counter, 1)
	if current <= rl.limit {
		return true
	}
	atomic.AddInt64(&rl.counter, -1)
	return false
}

func (rl *RateLimiter) CurrentCount() int64 {
	return atomic.LoadInt64(&rl.counter)
}

func (rl *RateLimiter) resetLoop() {
	ticker := time.NewTicker(rl.interval)
	defer ticker.Stop()

	for range ticker.C {
		atomic.StoreInt64(&rl.counter, 0)
		fmt.Println(">>> Rate limiter сброшен <<<")
	}
}

func main() {
	rl := NewRateLimiter(5, 1*time.Second)

	for i := 0; i < 20; i++ {
		if rl.Allow() {
			fmt.Printf("Запрос %d:  разрешен\n", i+1)
		} else {
			fmt.Printf("Запрос %d:  отклонен \n", i+1)
		}
		time.Sleep(200 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
}
