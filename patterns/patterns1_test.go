package patterns

import (
	"log"
	"testing"
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

func TestPatterns1(t *testing.T) {
	c := gen(
		item{
			price:    8,
			category: "shirt",
			discount: 0,
		},
		item{20, "shoe", 0.05},
		item{24, "shoe", 0.5},
		item{4, "drink", 0},
	)

	out := discount(c)
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
