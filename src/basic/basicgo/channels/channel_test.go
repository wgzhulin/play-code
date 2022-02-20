package channels

import (
	"log"
	"testing"
)

func TestChannel(t *testing.T) {
	done := make(chan struct{})
	go func() {
		log.Println("done")
		done <- struct{}{}
	}()

	<-done
}

func Test_Channel_Int(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 2
	log.Println(<-ch)
	close(ch)
}

func Test_Channel_Squarer(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i <= 100; i++ {
			naturals <- i
		}
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
	}()

	for x := range squares {
		log.Println(x)
	}
}

func Test_Channel_SingleDirectChannel(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	counter := func(in chan<- int) {
		for i := 0; i <= 100; i++ {
			in <- i
		}
	}

	squarer := func(out chan<- int, in <-chan int) {
		for x := range in {
			out <- x * x
		}
	}

	printer := func(in <-chan int) {
		for v := range in {
			log.Println(v)
		}
	}

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
