// **Задача 5.**
// Напишите функцию `Merge(ch1, ch2 <-chan int) <-chan int`, которая объединяет два канала в один.
// Горутины-писатели должны завершиться корректно,
// а выходной канал — закрыться после исчерпания обоих входных. Используйте `sync.WaitGroup`.
package main

import "sync"

func Merge(ch1, ch2 <-chan int) <-chan int {
	ch12 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for val := range ch1 {
			ch12 <- val
		}
	}()
	go func() {
		defer wg.Done()
		for val := range ch2 {
			ch12 <- val
		}
	}()
	go func() {
		wg.Wait()
		close(ch12)
	}()

	return ch12
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()
	go func() {
		ch2 <- 5
		ch2 <- 6
		ch2 <- 7
		close(ch2)
	}()
	merged := Merge(ch1, ch2)
	for v := range merged {
		println(v)
	}
}
