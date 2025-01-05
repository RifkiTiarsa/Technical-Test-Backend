package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

type BankAccount struct {
	Balance float64
	Mutex   sync.Mutex
}

func main() {
	account := BankAccount{Balance: 0}

	for i := 1; i <= 10; i++ {
		wg.Add(2)
		go account.topUp(3000)
		go account.withDraw(1000)
	}

	wg.Wait()
	fmt.Println("Your current balance : ", account.Balance)
}

func (account *BankAccount) topUp(amount float64) {
	defer wg.Done()
	defer account.Mutex.Unlock()
	account.Mutex.Lock()
	account.Balance += amount
}

func (account *BankAccount) withDraw(amount float64) {
	defer account.Mutex.Unlock()
	defer wg.Done()
	account.Mutex.Lock()
	if account.Balance >= amount {
		account.Balance -= amount
	}
}
