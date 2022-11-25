package ping_pong

import (
	"log"
	"math/rand"
	"time"
)

// YouTube: @imrenagi
// Dengan solusi sebelumnya, tambahkan fitur:
//   - Goroutine untuk referree/wasit
//   - Wasit akan mangambil bola ketika salah satu pemain
//     tidak dapat mengembalikan bola ke lawan
//   - Permainan selesai, wasit menentukan pemenang.
type ballReferree struct {
	hits       int
	lastPlayer string
}

func referree(table chan *ballReferree, done chan *ballReferree) {
	// new(ball) <=> &ball{}
	table <- &ballReferree{}

	for {
		select {
		case ball := <-done:
			log.Println("winner is", ball.lastPlayer)
			return
		}
	}
}

func playerReferree(name string, table chan *ballReferree, done chan *ballReferree) {
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
