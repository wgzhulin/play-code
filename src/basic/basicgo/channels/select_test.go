package channels

import (
	"log"
	"testing"
)

func TestSelectChannel(t *testing.T) {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			log.Println(x)
		case ch <- i:
		}
	}
}
