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

func TestPatternsPipeline(t *testing.T) {
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

	out := Discount(c)
	for processes := range out {
		// // @DonaldFeury
		// fmt.Println("Category:", processes.category, "Price:", processes.price)

		// t.Log("Category:", processes.category+",", "Price: $", processes.price) -> no debug test
		log.Print("Category:", processes.Category+",", " Price: $", processes.Price)
		if processes.Discount > 0 {
			log.Print("\t and the discount is $", processes.Discount)
		}
	}
}
