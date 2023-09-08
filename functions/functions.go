package functions

import (
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// function with named parameters
func swapVals(a, b int) (x, y int) {
	fmt.Println("Swapping values")
	y, x = a, b
	return
}

// variadic function
func Min(b ...int) int {
	if len(b) == 0 {
		return 0
	}
	min := b[0]
	for _, v := range b {
		if v < min {
			min = v
		}
	}
	return min
}

// variadic function passing its params
func passParams(args ...int) {
	fmt.Println("Now i m gonna pass variadic args to Min\u2BB0")
	fmt.Printf("The min = %d\n", Min(args...))
}

// use of variadic interface{}
func typecheck(values ...interface{}) {
	fmt.Println(">>>>Typecheck:")
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
	fmt.Println("Typecheck<<<")
}
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

// Options
type Options struct {
	Param1 int
	Param2 string
	Param3 bool
}

func doStuffWithOptions(a, b int, opts Options) {
	fmt.Printf("a=%d b=%d\n", a, b)
	v := reflect.ValueOf(opts)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s, Value: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())

	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
func factorial(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	} else if n == 1 || n == 2 {
		return big.NewInt(int64(n))
	}
	result := big.NewInt(1)
	result.Mul(factorial(n-1), factorial(n-2))
	return result
}

func fibonacci_closure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func isAscii(c rune) bool {
	if c > 255 {
		return false
	}
	return true
}

const LIM = 41

var cache [LIM]uint64

// fibo with memo
func fibo_with_memo(n int) (res uint64) {
	if cache[n] != 0 {
		res = cache[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibo_with_memo(n-1) + fibo_with_memo(n-2)
	}
	cache[n] = res
	return
}

func Run() {
	fmt.Println(swapVals(7, 96))
	fmt.Printf("The min = %d\n", Min(7, 5, 8, 8, 3, 99))
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply: ", *reply)
	fmt.Println(n)
	typecheck(1, 5, "aha")
	passParams(1, -5, 3)
	doStuffWithOptions(1, 1, Options{15, "ABC", true})
	for i := 1; i <= 10; i++ {
		fmt.Printf("fibonacci(%d) = %d\n", i, fibonacci(i))
	}
	fmt.Println("Factorial of 22 =", factorial(22))
	fmt.Println(strings.IndexFunc("\u2BCC\u2BCCA", isAscii))
	fmt.Println("\u2BCC\u2BCDA")
	fmt.Println("Fibonacci with closures⮰")
	var a = fibonacci_closure()
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Print(a(), " ")
	fmt.Println()
	var where = func() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
	}
	where()
	where()
	where()

	fmt.Println("Fibo with memoization")
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result := fibo_with_memo(i)
		fmt.Printf("fiboMemo(%d) is: %d\n", i, result)
	}
	end := time.Now()
	fmt.Printf("fibo with memo took this amount of time: %s\n", end.Sub(start))
	start = time.Now()
	for i := 0; i < LIM; i++ {
		result := fibonacci(i)
		fmt.Printf("fiboMemo(%d) is: %d\n", i, result)
	}
	end = time.Now()
	fmt.Printf("usual fibo took this amount of time: %s\n", end.Sub(start))
}
