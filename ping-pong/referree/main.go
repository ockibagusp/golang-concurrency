package main

import (
	"log"
	"math/rand"
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
	done := make(chan *ball)

	go player("imre", table, done)
	go player("ocki", table, done)

	// // new(ball) <=> &ball{}
	// table <- &ball{}
	// // // time.Sleep(1 * time.Second)
	// // 2022/11/22 22:53:07 ocki hits the ball 1
	// // 2022/11/22 22:53:07 imre hits the ball 2
	// time.Sleep(1 * time.Second)
	// <-table
	referree(table, done)
}

type ball struct {
	hits       int
	lastPlayer string
}

func referree(table chan *ball, done chan *ball) {
	// new(ball) <=> &ball{}
	table <- &ball{}

	for {
		select {
		case ball := <-done:
			log.Println("winner is", ball.lastPlayer)
			return
		}
	}
}

func player(name string, table chan *ball, done chan *ball) {
	for {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		select {
		case ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				log.Println(name, "drop the ball")
				done <- ball
				return
			}

			ball.hits++
			ball.lastPlayer = name
			log.Println(name, "hits the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball
		case <-time.After(2 * time.Second):
			return
		}
	}
}
