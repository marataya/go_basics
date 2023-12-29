package average_of_numbers

import (
	"fmt"
	"os"
)

func Run() {
	//f := squares()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	var sum float64
	var n int
	for {
		var val float64
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Println(os.Stderr, "no values")
	}

	fmt.Println("The average is ", sum / float64(n))
}

