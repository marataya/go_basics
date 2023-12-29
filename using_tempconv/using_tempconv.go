package using_tempconv

import (
	"fmt"
	"go_basics/tempconv"
)

func Run() {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToF(tempconv.FreezingC))
}
