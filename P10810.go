// Slice
// https://www.acmicpc.net/problem/10810

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
	basket := make([]int, N+1)
	for n := 0; n < M; n++ {
		i, j, k := scanInt(), scanInt(), scanInt()
		for p := i; p <= j; p++ {
			basket[p] = k
		}
	}
	for i := 1; i < N; i++ {
		fmt.Print(basket[i], " ")
	}
	fmt.Print(basket[N])
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
