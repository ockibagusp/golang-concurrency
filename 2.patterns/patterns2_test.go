package patterns

import (
	"log"
	"sync"
	"testing"
	"time"
)

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

type item2 struct {
	price    float32
	category string
	discount float32
}

func gen2(items ...item2) <-chan item2 {
	// out := make(chan item, 4)
	out := make(chan item2, len(items))
	for _, i := range items {
		out <- i
	}
	close(out)
	return out
}

func discount2(items <-chan item2) <-chan item2 {
	out := make(chan item2)
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

func fanIn(channels ...<-chan item2) <-chan item2 {
	var wg sync.WaitGroup
	out := make(chan item2)
	output := func(c <-chan item2) {
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

func TestPatterns2(t *testing.T) {
	c := gen2(
		item2{
			price:    8,
			category: "shirt",
			discount: 0,
		},
		item2{20, "shoe", 0.05},
		item2{24, "shoe", 0.5},
		item2{4, "drink", 0},
	)

	c1 := discount2(c)
	c2 := discount2(c)
	out := fanIn(c1, c2)
	for processes := range out {
		// // @DonaldFeury
		// fmt.Println("Category:", processes.category, "Price:", processes.price)

		// t.Log("Category:", processes.category+",", "Price: $", processes.price) -> no debug test
		log.Print("Category:", processes.category+",", " Price: $", processes.price)
		if processes.discount > 0 {
			log.Print("\t and the discount is $", processes.discount)
		}
	}
}
