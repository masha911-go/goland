// **Задача 10.**
// Создайте структуру `OnceFlag` с методом `Do(f func())`, который гарантирует, что `f` будет вызвана **ровно один раз**,
// даже если `Do` вызывают несколько горутин одновременно. Используйте `sync.Mutex` (не используйте `sync.Once`!).
package main

import (
	"fmt"
	"sync"
)

type OnceFlag interface {
	Do(f func())
}

type Flag struct {
	count int
	flag  bool
}

func (f *Flag) f(s int, mutex *sync.Mutex) {
	mutex.Lock()
	mutex.Unlock()
	f.count += s
}

func (f *Flag) Do(s int, t bool, mutex *sync.Mutex) {
	mutex.Lock()
	mutex.Unlock()
	if f.flag == false {
		f.f(s, mutex)
	} else {
		return
	}

	f.flag = t
}

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	flak := Flag{count: 0, flag: false}
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			flak.Do(1, true, mu)
		}()
	}
	wg.Wait()
	fmt.Printf("%d\n", flak.count)
}
