package bank

import (
	"sync"
	"testing"
)

// Github: @ProgrammerZamanNow
// Youtube: @mattkdvb5154

type BankAccountMutex struct {
	Mutex     sync.Mutex
	WaitGroup sync.WaitGroup
	Balance   int
}

func (account *BankAccountMutex) SumBalance(amount int) {
	account.Mutex.Lock()
	account.Balance += amount
	account.Mutex.Unlock()
	account.WaitGroup.Done()
}

func (account *BankAccountMutex) GetBalance() int {
	account.Mutex.Lock()
	balance := account.Balance
	account.Mutex.Unlock()
	return balance
}

func TestBankAccountMutex(t *testing.T) {
	account := BankAccountMutex{}

	// add: +10.000
	account.WaitGroup.Add(1)
	go account.SumBalance(10000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalance(); getbalance != 10000 {
		t.Errorf("account.GetBalance == 10000; want: %d", getbalance)
	}

	// reduce: -5000
	account.WaitGroup.Add(1)
	go account.SumBalance(-5000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalance(); getbalance != 5000 {
		t.Errorf("account.GetBalance == 5000; want: %d", getbalance)
	}

	// reduce: -5000
	account.WaitGroup.Add(1)
	go account.SumBalance(-5000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalance(); getbalance != 0 {
		t.Errorf("account.GetBalance == 0; want: %d", getbalance)
	}

	// add: 1 * 100 => 100
	for i := 0; i < 100; i++ {
		account.WaitGroup.Add(1)
		go func() {
			account.SumBalance(1)
		}()
	}
	account.WaitGroup.Wait()

	if getbalance := account.GetBalance(); getbalance != 100 {
		t.Errorf("account.GetBalance == 100; want: %d", getbalance)
	}
}
