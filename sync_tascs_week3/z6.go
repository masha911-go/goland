// 6. Атомарный генератор ID
// Реализуйте генератор уникальных идентификаторов, который гарантированно возвращает
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type IDGenerator struct {
	counter uint64
}

func (g *IDGenerator) Next() uint64 {
	return atomic.AddUint64(&g.counter, 1)
}

func main() {
	var gen IDGenerator
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(gen.Next())
		}()
	}

	wg.Wait()
}
