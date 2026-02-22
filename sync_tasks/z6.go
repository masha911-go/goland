// **Задача 6.**
// Создайте «ограниченный семафор» с помощью буферизованного канала ёмкостью 3. Запустите 10 горутин, каждая из которых
// «занимает слот», ждёт 100 мс, затем освобождает его. Убедитесь, что одновременно работает не более 3 горутин.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	semaphore := make(chan struct{}, 3)
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}
			fmt.Println("занимает слот")
			time.Sleep(100 * time.Millisecond)
			<-semaphore
			fmt.Println("освобождает слот")
		}()
	}
	wg.Wait()
}
