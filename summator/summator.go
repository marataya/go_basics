package summator

import (
	"fmt"
)

func Run() {
	var n int
	var result []int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		result = append(result, a+b)
	}
	for i := 0; i < n; i++ {
		fmt.Println(result[i])
	}

}
