// Greedy
// https://www.acmicpc.net/problem/25635

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

	N := scanInt()
	slice := make([]int, 0)
	sum, max := 0, 0
	
	for i := 0; i < N; i++ {
		val := scanInt()
		slice = append(slice, val)
		sum += val
		if max < val {
			max = val
		}
	}

	valid := 0
	if sum % 2 == 0 {
		valid = sum / 2
	} else {
		valid = sum / 2 + 1
	}

	if valid < max {
		fmt.Println((sum - max) * 2 + 1)
		os.Exit(0);
	}

	fmt.Println(sum)
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
