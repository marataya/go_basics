package wc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Linux word counter utility WC

func Run() {
	var glc, gwc, gcc int
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		defer file.Close()

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}
		glc += lc
		gwc += wc
		gcc += cc
		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
	}
	fmt.Printf("%7d %7d %7d total\n", glc, gwc, gcc)
}
