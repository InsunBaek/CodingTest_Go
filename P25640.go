// IO
// https://www.acmicpc.net/problem/25640

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

	Jinho := scanString()
	N := scanInt()
	dup := 0

	for i := 0; i < N; i++ {
		if Jinho == scanString() {
			dup += 1
		}
	}

	fmt.Fprintln(bw, dup)
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}