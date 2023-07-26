package main

import (
	"fmt"
)

// make a map of maps that counts number of occurences of each name in a slice and places it map of string to int and
// then in turn in map of fist char of the string (rune) to names
func getNamesCount(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar := rune(name[0])
		_, ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}

// check if two maps are equal
func equal[T, U comparable](x, y map[T]U) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(getNamesCount([]string{"Connor", "Jim", "Joe", "Connor"}))
	mapA := map[string]int{
		"aaa": 1,
		"bbb": 2,
	}
	mapB := map[string]int{
		"aaa": 1,
		"bbb": 2,
	}
	fmt.Println(equal(mapA, mapB))

}
