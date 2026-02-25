// 4. Thread-safe булев флаг
// Реализуйте булев флаг, который можно безопасно устанавливать и проверять из разных горутин, используя только atomic.
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type BoolFlag int32

func (f *BoolFlag) SetTrue() {
	atomic.StoreInt32((*int32)(f), 1)
}

func (f *BoolFlag) SetFalse() {
	atomic.StoreInt32((*int32)(f), 0)
}

func (f *BoolFlag) IsSet() bool {
	return atomic.LoadInt32((*int32)(f)) == 1
}

func main() {
	var flag BoolFlag

	for i := 0; i < 3; i++ {
		go func(id int) {
			ticker := time.NewTicker(500 * time.Millisecond)
			for range ticker.C {
				if flag.IsSet() {
					fmt.Printf("%d: TRUE\n", id)
				} else {
					fmt.Printf("%d: false\n", id)
				}
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	flag.SetTrue()
	fmt.Println("flag = true")

	time.Sleep(2 * time.Second)
	flag.SetFalse()
	fmt.Println("flag = false")

	time.Sleep(2 * time.Second)
}
