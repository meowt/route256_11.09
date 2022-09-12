package main

import (
	"bufio"
	"bytes"
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
		var str string
		fmt.Fscan(in, &str)
		var buffer bytes.Buffer
		fmt.Println(fmt.Printf("%c", str[i]))
	LABEL:
		for i := 0; i < len(str); {
			if str[i] == 0 && str[i+1] == 0 {
				buffer.WriteString("a")
				i += 2
				goto LABEL
			}
			if str[i] == 1 && str[i+1] == 1 {
				buffer.WriteString("d")
				i += 2
				goto LABEL
			}
			if str[i] == 1 && str[i+1] == 0 {
				if str[i+2] == 0 {
					buffer.WriteString("b")
					i += 3
					goto LABEL
				}
				if str[i+2] == 1 {
					buffer.WriteString("c")
					i += 3
					goto LABEL
				}
			}
		}
		fmt.Fprintln(out, buffer.String())
	}
}
