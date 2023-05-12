// Simulation
// https://www.acmicpc.net/problem/25642

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

	YongTae, Yujin := scanInt(), scanInt()
	fmt.Fprintln(bw, Simulation(YongTae, Yujin))
}

func Simulation(YongTae, Yujin int) string {
	for  {
		Yujin += YongTae
		if Yujin >= 5 {
			return "yt"
		}

		YongTae += Yujin
		if YongTae >= 5 {
			return "yj"
		}
	}
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
