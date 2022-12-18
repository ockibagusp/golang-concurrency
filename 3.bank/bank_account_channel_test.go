package bank

import (
	"sync"
	"testing"
)

// Github: @ProgrammerZamanNow
// Youtube: @mattkdvb5154
// https://golangbot.com/mutex/

type BankAccountChannel struct {
	sync.WaitGroup
	Balance int
}

func (account *BankAccountChannel) SumBalanceChannel(channel chan bool, amount int) {
	// 1. defer
	defer account.WaitGroup.Done()

	channel <- true
	account.Balance += amount
	<-channel
	// 2. or...
	// account.WaitGroup.Done()
}

func (account *BankAccountChannel) GetBalanceChannel(channel chan bool) int {
	channel <- true
	balance := account.Balance
	<-channel
	return balance
}

func TestBankAccountChannel(t *testing.T) {
	// channel := make(chan bool)
	// // fatal error: all goroutines are asleep - deadlock!
	channel := make(chan bool, 1)
	// PASS

	account := BankAccountChannel{}

	// add: +10.000
	account.WaitGroup.Add(1)
	go account.SumBalanceChannel(channel, 10000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalanceChannel(channel); getbalance != 10000 {
		t.Errorf("account.GetBalance == 10000; want: %d", getbalance)
	}

	// reduce: -5000
	account.WaitGroup.Add(1)
	go account.SumBalanceChannel(channel, -5000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalanceChannel(channel); getbalance != 5000 {
		t.Errorf("account.GetBalance == 5000; want: %d", getbalance)
	}

	// reduce: -5000
	account.WaitGroup.Add(1)
	go account.SumBalanceChannel(channel, -5000)
	account.WaitGroup.Wait()

	if getbalance := account.GetBalanceChannel(channel); getbalance != 0 {
		t.Errorf("account.GetBalance == 0; want: %d", getbalance)
	}

	// add: 1 * 100 => 100
	for i := 0; i < 100; i++ {
		account.WaitGroup.Add(1)
		go func() {
			account.SumBalanceChannel(channel, 1)
		}()
	}
	account.WaitGroup.Wait()

	if getbalance := account.GetBalanceChannel(channel); getbalance != 100 {
		t.Errorf("account.GetBalance == 100; want: %d", getbalance)
	}
}
