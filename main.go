package main

import (
	"fmt"
	"time"
)

type Flags int

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
)

func main() {
	var td time.Duration
	t0 := time.Now()
	c := 0
	for i := 0; i < 1000; i++ {
		c++
	}
	t1 := time.Now()
	td = t1.Sub(t0)
	fmt.Println(td)
}
