package main

import (
	"log"
	"time"
)

// YouTube: @imrenagi
// Dengan solusi sebelumnya, tambahkan fitur:
// - Goroutine untuk referree/wasit
// - Wasit akan mangambil bola ketika salah satu pemain
// 	tidak dapat mengembalikan bola ke lawan
// - Permainan selesai, wasit menentukan pemenang.

func main() {
	table := make(chan *ball)

	go player("imre", table)
	go player("ocki", table)

	// new(ball) <=> &ball{}
	table <- &ball{}
	// // time.Sleep(1 * time.Second)
	// 2022/11/22 22:53:07 ocki hits the ball 1
	// 2022/11/22 22:53:07 imre hits the ball 2
	time.Sleep(1 * time.Second)
	<-table
}

type ball struct {
	hits int
}

func player(name string, table chan *ball) {
	for {
		ball := <-table
		ball.hits++
		log.Println(name, "hits the ball", ball.hits)
		time.Sleep(50 * time.Millisecond)
		table <- ball
	}
}
