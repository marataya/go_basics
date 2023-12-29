package slices

import (
	"bytes"
	"fmt"
)

func sumArray(arr []float64) (res float64) {
	for _, v := range arr {
		res += v
	}
	return
}

func Append(slice, data []byte) []byte {
	l := len(slice)
	//extending capacity when new length is above cap
	if l+len(data) > cap(slice) {
		slice = append(slice, make([]byte, l+len(data))...)
	}
	//if l+len(data) > cap(slice) {
	//	newSlice := make([]byte, 2*(l+len(data)))
	//	copy(newSlice, slice)
	//	slice = newSlice
	//}
	//reslicing
	slice = slice[0 : l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

func Split(s string, pos int) []string {
	s1, s2 := s[:pos], s[pos:]
	return []string{s1, s2}
}

func StringReverse(s string) (res string) {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		res += string(runes[i])
	}
	return
}
func RunEx1() {
	var b = []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[:2])
	fmt.Println(b[2:])
	fmt.Println(b[:])
	fmt.Println("using bytes.Join\t", bytes.Join([][]byte{b[:2], b[2:]}, []byte("")))
	x := append(b[:2], b[2:]...)
	fmt.Println("using append on slices\t", x)
	fmt.Println("Append []byte to slice\u2bae")
	fmt.Println(Append(b, []byte{'x', 'x', 'x'}))

	items := [...]int{10, 20, 30, 40, 50}
	for idx, _ := range items {
		items[idx] *= 2
	}
	fmt.Println(items)
	arr := []float64{6, 7, 8.8, 19.7, -6.5}
	fmt.Println("Sum of", arr, "=", sumArray(arr))
	s := "Motherfucker"
	fmt.Println("array s split at pos 5:", Split(s, 6))
	fmt.Println("Reverse string:", StringReverse("ФЫВА"))
}
