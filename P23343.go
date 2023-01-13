// Implementation
// https://www.acmicpc.net/problem/23343

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	x, xF := specialScan()
	y, yF := specialScan()

	if xF || yF {
		fmt.Fprintln(bw, "NaN")
		return
	}
	fmt.Fprintln(bw, x-y)
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn()
}

func specialScan() (int, bool) {
	sc.Scan()
	v, err := strconv.Atoi(sc.Text())
	if err != nil {
		return 0, true
	}
	return v, false
}
