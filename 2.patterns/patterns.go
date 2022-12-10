package patterns

import (
	"sync"
	"time"
)

// YouTube: @DonaldFeury

/*
Pipeline
--------

routine -> routine -> routine
*/

type item struct {
	price    float32
	category string
	discount float32
}

func gen(items ...item) <-chan item {
	// out := make(chan item, 4)
	out := make(chan item, len(items))
	for _, i := range items {
		out <- i
	}
	close(out)
	return out
}

func discount(items <-chan item) <-chan item {
	out := make(chan item)
	go func() {
		defer close(out)
		for i := range items {
			// We have a sale going on
			// Shoes are half off!
			if i.category == "shoe" {
				// // @DonaldFeury
				// i.price /= 2

				// 50  / 100 = 0 (int) x
				// 50.0 / 100.0 = 0.5 (float32) v
				i.discount = i.price * i.discount
				i.price = i.price - i.discount // 50%
				// // i.price = i.price - (i.price * i.discount)
			}
			out <- i
		}
	}()
	return out
}

// YouTube: @DonaldFeury

/*
Fan In/Fan Out
--------
          routine
        /		  \
routine				-> routine -> routine
		\		  /
		  routine
*/

func discountSleep(items <-chan item) <-chan item {
	out := make(chan item)
	go func() {
		defer close(out)
		for i := range items {
			time.Sleep(2 * time.Second)

			// We have a sale going on
			// Shoes are half off!
			if i.category == "shoe" {
				// // @DonaldFeury
				// i.price /= 2

				// 50  / 100 = 0 (int) x
				// 50.0 / 100.0 = 0.5 (float32) v
				i.discount = i.price * i.discount
				i.price = i.price - i.discount // 50%
				// // i.price = i.price - (i.price * i.discount)
			}
			out <- i
		}
	}()
	return out
}

func discountDone(done <-chan bool, items <-chan item) <-chan item {
	out := make(chan item)
	go func() {
		defer close(out)
		for i := range items {
			time.Sleep(2 * time.Second)

			// We have a sale going on
			// Shoes are half off!
			if i.category == "shoe" {
				// // @DonaldFeury
				// i.price /= 2

				// 50  / 100 = 0 (int) x
				// 50.0 / 100.0 = 0.5 (float32) v
				i.discount = i.price * i.discount
				i.price = i.price - i.discount // 50%
				// // i.price = i.price - (i.price * i.discount)
			}

			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}

func fanIn(channels ...<-chan item) <-chan item {
	var wg sync.WaitGroup
	out := make(chan item)
	output := func(c <-chan item) {
		defer wg.Done()
		for i := range c {
			out <- i
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fanInDone(done <-chan bool, channels ...<-chan item) <-chan item {
	var wg sync.WaitGroup
	out := make(chan item)
	output := func(c <-chan item) {
		defer wg.Done()
		for i := range c {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
