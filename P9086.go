// String
// https://www.acmicpc.net/problem/9086

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

	Tc := scanInt()
	for i := 0; i < Tc; i++ {
		str := scanString()
		fmt.Fprint(bw, string(str[0]))
		fmt.Fprint(bw, string(str[len(str) - 1]))
		fmt.Fprint(bw, "\n")
	}
}


func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}
