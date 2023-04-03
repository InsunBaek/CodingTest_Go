// String
// https://www.acmicpc.net/problem/27866

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
	
	S := scanString()
	fmt.Fprintln(bw, string(S[scanInt() - 1]))
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
