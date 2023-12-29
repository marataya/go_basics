package goroutines

import (
	"fmt"
	"time"
)

func generate(ch chan int) {
	for i := 2; i < 100; i++ {
		ch <- i
	}
}

// function removing divisible by prime
func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func Run1() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	time.Sleep(time.Second * 1)
}

func Run() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
