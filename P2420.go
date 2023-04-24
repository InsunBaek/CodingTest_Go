// Math
// https://www.acmicpc.net/problem/2420

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

	N, M := scanInt(), scanInt()
	sum := N - M
	if sum > 0 {
		fmt.Fprintln(bw, sum)
	} else {
		fmt.Fprintln(bw, -sum)
	}
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
