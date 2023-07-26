package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mult(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

// Function currying
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func main() {
	fmt.Println(aggregate(2, 3, 4, add))
	fmt.Println(aggregate(2, 3, 4, mult))
	squareFunc := selfMath(mult)
	doubleFunc := selfMath(add)
	fmt.Println("---------------------")
	fmt.Println("Function currying")
	fmt.Println(squareFunc(5))
	fmt.Println(doubleFunc(5))

	r := [...]string{5: "-1"}
	for _, el := range r {
		fmt.Printf("%q\t", el)
	}
}
