package bank

import (
	"sync"
	"testing"
)

// Github: @ProgrammerZamanNow
// Youtube: @mattkdvb5154

type BankAccountChannel struct {
	Balance int
}

func (account *BankAccountChannel) SumBalanceChannel(wg *sync.WaitGroup, channel chan bool, amount int) {
	channel <- true
	account.Balance += amount
	<-channel
	wg.Done()
}

func (account *BankAccountChannel) GetBalanceChannel(channel chan bool) int {
	channel <- true
	balance := account.Balance
	<-channel
	return balance
}

func TestBankAccountChannel(t *testing.T) {
	var wg sync.WaitGroup
	channel := make(chan bool, 1)

	account := BankAccountChannel{}

	// add: +10.000
	wg.Add(1)
	go account.SumBalanceChannel(&wg, channel, 10000)
	wg.Wait()

	if getbalance := account.GetBalanceChannel(channel); getbalance != 10000 {
		t.Errorf("account.GetBalance == 10000; want: %d", getbalance)
	}

	// reduce: -5000
	wg.Add(1)
	go account.SumBalanceChannel(&wg, channel, -5000)
	wg.Wait()

	if getbalance := account.GetBalanceChannel(channel); getbalance != 5000 {
		t.Errorf("account.GetBalance == 5000; want: %d", getbalance)
	}

	// reduce: -5000
	wg.Add(1)
	go account.SumBalanceChannel(&wg, channel, -5000)
	wg.Wait()

	if getbalance := account.GetBalanceChannel(channel); getbalance != 0 {
		t.Errorf("account.GetBalance == 0; want: %d", getbalance)
	}

	// add: 1 * 100 => 100
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			account.SumBalanceChannel(&wg, channel, 1)
		}()
	}
	wg.Wait()

	if getbalance := account.GetBalanceChannel(channel); getbalance != 100 {
		t.Errorf("account.GetBalance == 100; want: %d", getbalance)
	}
}
