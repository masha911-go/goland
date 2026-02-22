// **Задача 7.**
// Реализуйте конкурентно-безопасный логгер `SafeLogger`, который пишет строки в `os.Stdout`. Несколько горутин должны
// иметь возможность вызывать `Log(msg string)` без перехлёста строк.
package main

import (
	"fmt"
	"os"
	"sync"
)

type SafeLogger struct {
	sync.Mutex
}

func (s *SafeLogger) Log(msg string) {
	s.Lock()
	defer s.Unlock()
	fmt.Fprintln(os.Stdout, msg)
}

func main() {
	logger := &SafeLogger{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			logger.Log(fmt.Sprintf("hello %d", id))
		}(i)
	}
	wg.Wait()
}
