// ## Задачи:
// 1. Кэш с RW Mutex
// Реализуйте потокобезопасный кэш, где операции чтения происходят значительно чаще, чем записи.
// Используйте sync.RWMutex для оптимизации производительности.
package main

import (
	"fmt"
	"sync"
)

type Cash struct {
	mu   sync.RWMutex
	data map[int]string
}

func NewCash() *Cash {
	return &Cash{data: make(map[int]string)}
}

func (c *Cash) Get(key int) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *Cash) Set(key int, value string) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func main() {
	c := NewCash()
	wg := sync.WaitGroup{}
	wg.Add(2)
	c.Set(1, "1")
	c.Set(2, "2")
	if value, ok := c.Get(1); ok {
		fmt.Println(value)
	}
}
