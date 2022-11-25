package ping_pong

import (
	"log"
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
