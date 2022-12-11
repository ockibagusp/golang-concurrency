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

	c := Gen(
		Item{
			Price:    8,
			Category: "shirt",
			Discount: 0,
		},
		Item{20, "shoe", 0.05},
		Item{24, "shoe", 0.5},
		Item{4, "drink", 0},
	)

	c1 := DiscountDone(done, c)
	c2 := DiscountDone(done, c)
	out := FanInDone(done, c1, c2)
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
