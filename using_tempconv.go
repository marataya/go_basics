package main

import (
	"fmt"
	"go_basics/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToF(tempconv.FreezingC))
}
