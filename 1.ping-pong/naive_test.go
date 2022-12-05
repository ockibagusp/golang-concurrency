package ping_pong

import (
	"log"
	"testing"
	"time"
)

// YouTube: @imrenagi
// Dengan solusi sebelumnya, tambahkan fitur:
//   - Goroutine untuk referree/wasit
//   - Wasit akan mangambil bola ketika salah satu pemain
//     tidak dapat mengembalikan bola ke lawan
//   - Permainan selesai, wasit menentukan pemenang.
type ballNaive struct {
	hits int
}

func playerNaive(name string, table chan *ballNaive) {
	for {
		ball := <-table
		ball.hits++
		log.Println(name, "hits the ball", ball.hits)
		time.Sleep(50 * time.Millisecond)
		table <- ball
	}
}

func TestNaive(t *testing.T) {
	table := make(chan *ballNaive)

	go playerNaive("imre", table)
	go playerNaive("ocki", table)

	// new(ballNaive) <=> &ballNaive{}
	table <- &ballNaive{}
	// // time.Sleep(1 * time.Second)
	// 2022/11/22 22:53:07 ocki hits the ball 1
	// 2022/11/22 22:53:07 imre hits the ball 2
	time.Sleep(1 * time.Second)
	<-table
}
