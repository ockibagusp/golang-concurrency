package ping_pong

import (
	"testing"
	"time"
)

func TestNaive(t *testing.T) {
	table := make(chan *ballNaive)

	go playerNaive("imre", table)
	go playerNaive("ocki", table)

	// new(ball) <=> &ball{}
	table <- &ballNaive{}
	// // time.Sleep(1 * time.Second)
	// 2022/11/22 22:53:07 ocki hits the ball 1
	// 2022/11/22 22:53:07 imre hits the ball 2
	time.Sleep(1 * time.Second)
	<-table
}
