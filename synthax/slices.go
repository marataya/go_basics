package synthax

import "fmt"

func main() {
	a := make([]int, 3)

	fmt.Println("a: ", a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println("Appending 4 to b from a")
	b := append(a, 4)
	fmt.Println(b)
	fmt.Println("Address of b: ", &b[0])
	fmt.Println("Appending 5 to c from a")
	c := append(a, 5)
	fmt.Println(c)
	fmt.Println("Address of c: ", &c[0])

	//a := make([]int, 3, 5)
	//fmt.Println("Address of a: ", &a[0])
	//fmt.Println("a: ", a)
	//fmt.Println(len(a))
	//fmt.Println(cap(a))
	//fmt.Println("Appending 4 to b from a")
	//b := append(a, 4)
	//fmt.Println(b)
	//fmt.Println("Address of b: ", &b[0])
	//fmt.Println("Appending 5 to c from a")
	//c := append(a, 5)
	//fmt.Println(c)
	//fmt.Println("Address of c: ", &c[0])
	fmt.Println("===========================")

	i := make([]int, 3, 8)
	fmt.Println("Address of i: ", &i[0])
	fmt.Println("i: ", i)
	j := append(i, 4)
	fmt.Println(i, j)
	fmt.Println("Address of j: ", &j[0])
	g := append(i, 5)
	fmt.Println(i, j, g)
	fmt.Println("Address of g: ", &g[0])

}
