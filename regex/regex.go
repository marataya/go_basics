package regex

import (
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 64)
		return strconv.FormatFloat(v/2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found!")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
