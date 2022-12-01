package ping_pong

import (
	"log"
	"math/rand"
	"testing"
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
			if ball.lastPlayer == "" {
				log.Printf("winner is %s*", names[0])
			} else {
				log.Println("winner is", ball.lastPlayer)
			}

			return
		}
	}
}

// https://stackoverflow.com/questions/61189263/panic-runtime-error-index-out-of-range-0-with-length-0
var names []string = make([]string, 2)

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
			log.Println("player referree the time after...")
			return
		}
	}
}

func TestReferree2Persons(t *testing.T) {
	table := make(chan *ballReferree)
	done := make(chan *ballReferree)

	names = []string{"imre", "ocki"}
	// TODO: names -> "imre", "ocki" atau "ocki", "imre"
	go playerReferree(names[0], table, done)
	go playerReferree(names[1], table, done)

	// // new(ball) <=> &ball{}
	// table <- &ball{}
	// // // time.Sleep(1 * time.Second)
	// // 2022/11/22 22:53:07 ocki hits the ball 1
	// // 2022/11/22 22:53:07 imre hits the ball 2
	// time.Sleep(1 * time.Second)
	// <-table
	referree(table, done)
}
