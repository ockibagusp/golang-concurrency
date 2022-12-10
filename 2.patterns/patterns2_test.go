package patterns

import (
	"log"
	"testing"
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

func TestPatternsFanInFanOut(t *testing.T) {
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

	c1 := discount(c)
	c2 := discount(c)
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
