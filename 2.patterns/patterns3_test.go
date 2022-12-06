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

type item3 struct {
	price    float32
	category string
	discount float32
}

func gen3(items ...item3) <-chan item3 {
	// out := make(chan item, 4)
	out := make(chan item3, len(items))
	for _, i := range items {
		out <- i
	}
	close(out)
	return out
}

func discount3(done <-chan bool, items <-chan item3) <-chan item3 {
	out := make(chan item3)
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

func fanIn3(done <-chan bool, channels ...<-chan item3) <-chan item3 {
	var wg sync.WaitGroup
	out := make(chan item3)
	output := func(c <-chan item3) {
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

func TestPatterns3(t *testing.T) {
	done := make(chan bool)
	defer close(done)

	c := gen3(
		item3{
			price:    8,
			category: "shirt",
			discount: 0,
		},
		item3{20, "shoe", 0.05},
		item3{24, "shoe", 0.5},
		item3{4, "drink", 0},
	)

	c1 := discount3(done, c)
	c2 := discount3(done, c)
	out := fanIn3(done, c1, c2)
	// for processes := range out {
	// 	// // @DonaldFeury
	// 	// fmt.Println("Category:", processes.category, "Price:", processes.price)

	// 	// t.Log("Category:", processes.category+",", "Price: $", processes.price) -> no debug test
	// 	log.Print("Category:", processes.category+",", " Price: $", processes.price)
	// 	if processes.discount > 0 {
	// 		log.Print("\t and the discount is $", processes.discount)
	// 	}
	// }

	// @DonaldFeury
	log.Println(<-out)
	log.Println(<-out)

	log.Println(<-out)
}
