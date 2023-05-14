// BruteForce
// https://www.acmicpc.net/problem/25643

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

	n, m := scanInt(), scanInt()
	str := make([]string, 0)

	for i := 0; i < n; i++ {
		str = append(str, scanString())
	}

	for i := 0; i < n-1; i++ {
		str1 := str[i]
		str2 := str[i+1]
		able := false

		for j := 1; j <= m; j++ {
			if str1[m-j:] == str2[:j] {
				able = true
				break
			}
			if str1[:j] == str2[m-j:] {
				able = true
				break
			}
		}

		if !able {
			fmt.Println("0")
			os.Exit(0)
		}
	}

	fmt.Println("1")
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}