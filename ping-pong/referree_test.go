package ping_pong

import "testing"

func TestReferree(t *testing.T) {
	table := make(chan *ballReferree)
	done := make(chan *ballReferree)

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
