// Array
// https://www.acmicpc.net/problem/10811

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
	basket := make([]int, N)
	for i := 0; i < N; i++ {
		basket[i] = i + 1
	}

	for i := 0; i < M; i++ {
		from, to := scanInt(), scanInt()
		basket = flip(basket, from, to)
	}
	for i := 0; i < N; i++ {
		fmt.Fprint(bw, basket[i])
		if i != N - 1 {
			fmt.Fprint(bw, " ")
		}
	}
}

func flip(arr []int, from, to int) []int {
	rtnArr := make([]int, len(arr))
	copy(rtnArr, arr)
	subArr, k := arr[from - 1: to], 1
	for i := from - 1; i < to; i++ {
		rtnArr[i] = subArr[len(subArr) - k]
		k++
	}
	return rtnArr
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

