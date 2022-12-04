package patterns

import (
	"log"
	"testing"
)

// YouTube: @DonaldFeury

type item struct {
	price    int
	category string
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
				i.price /= 2
			}
			out <- i
		}
	}()
	return out
}

func TestPatterns1(t *testing.T) {
	c := gen(
		item{8, "shirt"},
		item{20, "shoe"},
		item{24, "shoe"},
		item{4, "drink"},
	)

	out := discount(c)
	for processes := range out {
		// fmt.Println("Category:", processes.category, "Price:", processes.price)
		log.Println("Category:", processes.category, "Price:", processes.price)
	}
}
