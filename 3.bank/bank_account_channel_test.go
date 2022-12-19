package bank

import (
	"sync"
	"testing"
)

// Github: @ProgrammerZamanNow
// Youtube: @mattkdvb5154
// https://golangbot.com/mutex/
// 		-> Mutex since this problem does not require any communication between the goroutines.
//		-> Hence mutex would be a natural fit.

type BankAccountChannel struct {
	channel chan bool
	*sync.WaitGroup
	Balance int
}

func (account *BankAccountChannel) SumBalanceChannel(amount int) {
	// 1. defer
	defer account.WaitGroup.Done()

	account.channel <- true
	account.Balance += amount
	<-account.channel
	// 2. or...
	// account.WaitGroup.Done()
}

func (account *BankAccountChannel) GetBalanceChannel() int {
	account.channel <- true
	balance := account.Balance
	<-account.channel
	return balance
}

func TestBankAccountChannel(t *testing.T) {
	// channel := make(chan bool)
	// // fatal error: all goroutines are asleep - deadlock!
	// channel := make(chan bool, 1)
	// PASS

	account := BankAccountChannel{
		channel:   make(chan bool, 1),
		WaitGroup: &sync.WaitGroup{},
	}

	// add: +10.000
	account.WaitGroup.Add(1)
	go account.SumBalanceChannel(10000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalanceChannel(); getbalance != 10000 {
		t.Errorf("account.GetBalance == 10000; want: %d", getbalance)
	}

	// reduce: -5000
	account.WaitGroup.Add(1)
	go account.SumBalanceChannel(-5000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalanceChannel(); getbalance != 5000 {
		t.Errorf("account.GetBalance == 5000; want: %d", getbalance)
	}

	// reduce: -5000
	account.WaitGroup.Add(1)
	go account.SumBalanceChannel(-5000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalanceChannel(); getbalance != 0 {
		t.Errorf("account.GetBalance == 0; want: %d", getbalance)
	}

	// add: 1 * 100 => 100
	for i := 0; i < 100; i++ {
		account.WaitGroup.Add(1)
		go func() {
			account.SumBalanceChannel(1)
		}()
	}
	account.WaitGroup.Wait()

	if getbalance := account.GetBalanceChannel(); getbalance != 100 {
		t.Errorf("account.GetBalance == 100; want: %d", getbalance)
	}
}
