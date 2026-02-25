// 3. Конфигурация с горячей перезагрузкой
// Разработайте систему управления конфигурацией, которая позволяет:
// Часто читать параметры конфигурации (много горутин)
// Редко полностью перезагружать конфигурацию (одна горутина)
package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	Port int
	Host string
}

var (
	config Config
	mu     sync.RWMutex
)

func main() {
	mu.Lock()
	config = Config{Port: 8080, Host: "localhost"}
	mu.Unlock()

	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				mu.RLock()
				c := config
				mu.RUnlock()
				fmt.Printf("Reader %d: %+v\n", id, c)
				time.Sleep(500 * time.Millisecond)
			}
		}(i)
	}

	go func() {
		time.Sleep(2 * time.Second)

		mu.Lock()
		config = Config{Port: 9090, Host: "192.168.1.1"}
		mu.Unlock()
		fmt.Println(">>> Updated to 9090 <<<")

		time.Sleep(2 * time.Second)

		mu.Lock()
		config = Config{Port: 3000, Host: "example.com"}
		mu.Unlock()
		fmt.Println(">>> Updated to 3000 <<<")
	}()

	time.Sleep(10 * time.Second)
}
