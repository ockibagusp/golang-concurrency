package main

import (
	"fmt"
	"log"

	"github.com/ockibagusp/golang-concurrency/2.patterns/patterns"
)

func main() {
	// func Gen(items ...Item) <-chan Item {...}
	items := []patterns.Item{
		{
			Price:    8,
			Category: "shirt",
			Discount: 0,
		},
		{20, "shoe", 0.05},
		{24, "shoe", 0.5},
		{4, "drink", 0},
	}

	// 1. pipeline
	fmt.Println("1. pipeline")

	// func Gen(items ...Item) <-chan Item {...}
	//          ---------
	c := patterns.Gen(items...)

	out := patterns.Discount(c)
	for processes := range out {
		// // @DonaldFeury
		// fmt.Println("Category:", processes.category, "Price:", processes.price)

		// t.Log("Category:", processes.category+",", "Price: $", processes.price) -> no debug test
		log.Print("1. Category:", processes.Category+",", " Price: $", processes.Price)
		if processes.Discount > 0 {
			log.Print("\t and the discount is $", processes.Discount)
		}
	}

	// 2. Fan In/Fan Out
	fmt.Println()
	fmt.Println("2. Fan In/Fan Out")

	c = patterns.Gen(items...)

	c1 := patterns.Discount(c)
	c2 := patterns.Discount(c)
	out = patterns.FanIn(c1, c2)
	for processes := range out {
		// // @DonaldFeury
		// fmt.Println("Category:", processes.category, "Price:", processes.price)

		// t.Log("Category:", processes.category+",", "Price: $", processes.price) -> no debug test
		log.Print("2. Category:", processes.Category+",", " Price: $", processes.Price)
		if processes.Discount > 0 {
			log.Print("\t and the discount is $", processes.Discount)
		}
	}

	// 3. Fan In/Fan Out 2
	fmt.Println()
	fmt.Println("3. Fan In/Fan Out")

	done := make(chan bool)
	defer close(done)

	c = patterns.Gen(items...)
	c1 = patterns.DiscountDone(done, c)
	c2 = patterns.DiscountDone(done, c)
	out = patterns.FanInDone(done, c1, c2)
	// for processes := range out {
	// 	// @DonaldFeury
	//	fmt.Println("Category:", processes.Category, "Price:", processes.Price)

	// 	// t.Log("Category:", processes.Category+",", "Price: $", processes.Price) -> no debug test
	// 	log.Print("3. Category:", processes.Category+",", " Price: $", processes.Price)
	// 	if processes.Discount > 0 {
	// 		log.Print("\t and the discount is $", processes.Discount)
	// 	}
	// }

	// @DonaldFeury
	log.Println(<-out)
	log.Println(<-out)

	log.Println(<-out)
}
