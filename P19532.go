// BruteForce
// https://www.acmicpc.net/problem/19532

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

	a, b, c, d, e, f := scanInt(), scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
	
	for x := -999; x <= 999; x++ {
		for y := -999; y <= 999; y++ {
			if a * x + b * y == c && d * x + e * y == f {
				fmt.Println(x, y)
				break
			}
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
