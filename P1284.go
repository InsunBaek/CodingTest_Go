// Math
// https://www.acmicpc.net/problem/1284

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
		banner := scanInt()
		sum := 1

		if banner == 0 {
			break
		}

		for banner != 0 {
			unit := banner % 10
			banner /= 10
			if unit == 0 {
				sum += 4
			} else if unit == 1 {
				sum += 2
			} else {
				sum += 3
			}
			sum += 1
		}
		fmt.Fprintln(bw, sum)
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
