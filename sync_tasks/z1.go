// **Задача 1.**
// Создайте глобальную переменную `counter int`. Запустите 100 горутин, каждая из которых увеличивает `counter` на 1.
// Используйте `sync.Mutex` для защиты доступа. Убедитесь, что результат равен 100.
package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup
var counter int

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
