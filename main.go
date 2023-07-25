package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(reflect.TypeOf(arr))
}
