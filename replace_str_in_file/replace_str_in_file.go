package replace_str_in_file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	if len(os.Args) < 3 {
		fmt.Println(os.Stderr, "not enough arguments")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}

}

