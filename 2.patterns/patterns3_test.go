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

func TestPatternsFanInFanOut2(t *testing.T) {
	done := make(chan bool)
	defer close(done)

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

	c1 := discountDone(done, c)
	c2 := discountDone(done, c)
	out := fanInDone(done, c1, c2)
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
