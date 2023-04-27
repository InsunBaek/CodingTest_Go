// Math
// https://www.acmicpc.net/problem/2959

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	slice := make([]int, 0)
	for i := 0; i < 4; i++ {
		slice = append(slice, scanInt())
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	var minA, minB int
	if slice[0] < slice[1] {
		minA = slice[0]
	} else {
		minA = slice[1]
	}

	if slice[2] < slice[3] {
		minB = slice[2]
	} else {
		minB = slice[3]
	}

	fmt.Fprintln(bw, minA * minB)
}


func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
