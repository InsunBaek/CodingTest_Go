// Implementation
// https://www.acmicpc.net/problem/1544

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N := scanInt()
	src := make([]string, 0)
	var sb strings.Builder

	for i := 0; i < N; i++ {
		str := scanString()
		flag := false

		for j := 0; j < len(src); j++ {
			if strings.Contains(src[j], str) && (len(src[j])/2 == len(str)) {
				flag = true
				break
			}
		}
		if flag {
			continue
		} else {
			sb.WriteString(str)
			sb.WriteString(str)
			src = append(src, sb.String())
			sb.Reset()
		}
	}
	fmt.Println(len(src))
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
