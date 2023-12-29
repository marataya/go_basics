package anon_func

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func Run() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
