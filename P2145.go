// Math
// https://www.acmicpc.net/problem/2145

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	for {
		input := scanInt()
		if input == 0 {
			break
		}

		number := input

		for number / 10 > 0 {
			val := 0
			for number > 0 {
				val += (number % 10)
				number = number / 10
			}
			number = val
		}
		fmt.Fprintln(bw, number)
	}
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
