// Math
// https://www.acmicpc.net/problem/1075

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

	N, F := scanInt(), scanInt()
	N /= 100
	N *= 100
	for N % F != 0 {
		N++
	}
	N %= 100
	if N < 10 {
		fmt.Print("0")
		fmt.Println(N)
	} else {
		fmt.Println(N)
	}
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}