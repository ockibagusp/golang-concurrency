package multi

import (
	"log"
	"testing"
	"time"
)

func TestMulti(t *testing.T) {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {
		// 1. multi
		select {
		case m0 := <-chans[0]:
			log.Println("received:", m0)
		case m1 := <-chans[1]:
			log.Println("received:", m1)
		}
		// // ...
		// // 2022/12/20 10:46:14 received: 2
		// // 2022/12/20 10:46:14 received: 1
		// // 2022/12/20 10:46:15 received: 1
		// // 2022/12/20 10:46:16 received: 2
		// // 2022/12/20 10:46:16 received: 1

		// // 2. per satu-satu 1 dan 2
		// log.Println(<-chans[0])
		// log.Println(<-chans[1])

		// ...
		// // 2022/12/20 10:32:18 2
		// // 2022/12/20 10:32:18 1
		// // 2022/12/20 10:32:20 2
		// // 2022/12/20 10:32:20 1
		// // 2022/12/20 10:32:22 2
	}
}
