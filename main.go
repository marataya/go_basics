package main

import (
	"fmt"
)

func swapVals(a, b int) (x, y int) {
	y, x = a, b
	return
}
func Min(a int, b ...int) int {
	if len(b) == 0 {
		return a
	}
	min := a
	for _, v := range b {
		if v < min {
			min = v
		}
	}
	return min
}

// use of interface{}
func typecheck(values ...interface{}) {
	for _, value := range values {
		switch value.(type) {
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("default")
		}
	}
}
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	fmt.Println(swapVals(7, 96))
	fmt.Println(Min(7, 5, 8, 8, 3, 99))
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply: ", *reply)
	fmt.Println(n)
	typecheck(1, 5, "aha")
}
