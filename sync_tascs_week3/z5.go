// 5. Банковский счет с RW Mutex
// Создайте структуру банковского счета, где:
// Проверка баланса происходит очень часто
// Изменение баланса (пополнение/списание) происходит редко
package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	mu      sync.RWMutex
	balance float64
}

func (ba *BankAccount) GetBalance() float64 {
	ba.mu.RLock()
	defer ba.mu.RUnlock()
	return ba.balance
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.mu.Lock()
	ba.balance += amount
	ba.mu.Unlock()
}

func (ba *BankAccount) Withdraw(amount float64) {
	ba.mu.Lock()
	ba.balance -= amount
	ba.mu.Unlock()
}

func main() {
	acc := &BankAccount{balance: 1000}

	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				fmt.Printf("Читатель %d: %.2f\n", id, acc.GetBalance())
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	go func() {
		time.Sleep(1 * time.Second)
		acc.Deposit(500)
		fmt.Println("Пополнение +500")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		acc.Withdraw(300)
		fmt.Println("Снятие -300")
	}()

	time.Sleep(3 * time.Second)
}
