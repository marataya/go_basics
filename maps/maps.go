package maps

import (
	"fmt"
	"sort"
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

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16,
		"indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func Run() {
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
	//sorting maps
	fmt.Println("Unsorted map")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("Sorted map")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
