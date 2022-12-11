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

type Item struct {
	Price    float32
	Category string
	Discount float32
}

func Gen(items ...Item) <-chan Item {
	// out := make(chan item, 4)
	out := make(chan Item, len(items))
	for _, i := range items {
		out <- i
	}
	close(out)
	return out
}

func Discount(items <-chan Item) <-chan Item {
	out := make(chan Item)
	go func() {
		defer close(out)
		for i := range items {
			// We have a sale going on
			// Shoes are half off!
			if i.Category == "shoe" {
				// // @DonaldFeury
				// i.price /= 2

				// 50  / 100 = 0 (int) x
				// 50.0 / 100.0 = 0.5 (float32) v
				i.Discount = i.Price * i.Discount
				i.Price = i.Price - i.Discount // 50%
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

func DiscountSleep(items <-chan Item) <-chan Item {
	out := make(chan Item)
	go func() {
		defer close(out)
		for i := range items {
			time.Sleep(2 * time.Second)

			// We have a sale going on
			// Shoes are half off!
			if i.Category == "shoe" {
				// // @DonaldFeury
				// i.price /= 2

				// 50  / 100 = 0 (int) x
				// 50.0 / 100.0 = 0.5 (float32) v
				i.Discount = i.Price * i.Discount
				i.Price = i.Price - i.Discount // 50%
				// // i.price = i.price - (i.price * i.discount)
			}
			out <- i
		}
	}()
	return out
}

func DiscountDone(done <-chan bool, items <-chan Item) <-chan Item {
	out := make(chan Item)
	go func() {
		defer close(out)
		for i := range items {
			time.Sleep(2 * time.Second)

			// We have a sale going on
			// Shoes are half off!
			if i.Category == "shoe" {
				// // @DonaldFeury
				// i.price /= 2

				// 50  / 100 = 0 (int) x
				// 50.0 / 100.0 = 0.5 (float32) v
				i.Discount = i.Price * i.Discount
				i.Price = i.Price - i.Discount // 50%
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

func FanIn(channels ...<-chan Item) <-chan Item {
	var wg sync.WaitGroup
	out := make(chan Item)
	output := func(c <-chan Item) {
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

func FanInDone(done <-chan bool, channels ...<-chan Item) <-chan Item {
	var wg sync.WaitGroup
	out := make(chan Item)
	output := func(c <-chan Item) {
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
