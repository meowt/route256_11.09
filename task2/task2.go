package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var d, m, y int
		res := true
		fmt.Fscan(in, &d, &m, &y)
		if d > 29 {
			if m == 2 {
				res = false
			}

			if d > 30 {
				switch m {
				case 4:
					res = false
				case 6:
					res = false
				case 9:
					res = false
				case 11:
					res = false
				}
			}
		}
		if d == 29 {
			if !(y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m == 2 {
				res = false
			}
		}

		if res {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
