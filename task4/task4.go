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
		var testCounter2 int
		fmt.Fscan(in, &testCounter2)
		var slc []int
		var maxCounter1, maxCounter2, temp int
		for j := 0; j < testCounter2; j++ {
			var str int
			fmt.Fscan(in, &str)
			slc = append(slc, str)

			firstClient := slc[0]
			for i1 := 0; i1 < len(slc); i1++ {
				for i2 := i1 + 1; i2 < len(slc); i2++ {
					if slc[i1] == slc[i2] {
						temp++
					}
					if temp > maxCounter1 {
						firstClient = slc[i1]
						fmt.Println("first client is", firstClient)
						maxCounter1 = temp
						fmt.Println("max counter is", maxCounter1, "\n")
					}
				}
				temp = 0
			}
			for i1 := 0; i1 < len(slc); i1++ {
				if slc[i1] == firstClient {
					continue
				}
				for i2 := i1 + 1; i2 < len(slc); i2++ {
					if slc[i1] == slc[i2] {
						temp++
					}
					if temp > maxCounter1 {
						firstClient = slc[i1]
						maxCounter1 = temp
					}
				}
				temp = 0
			}
		}
		fmt.Fprintln(out, maxCounter1+maxCounter2)
	}
}
