// String
// https://www.acmicpc.net/problem/1652

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
	lines := make([]string, 0)
	
	for i := 0; i < N; i++ {
		lines = append(lines, scanString())
	}

	widthCnt := calc(lines, N)
	heightCnt := calc(reverseLines(lines, N), N)

	fmt.Println(widthCnt, heightCnt)
}

func calc(lines []string, N int) int {
	cnt := 0
	for row := 0; row < len(lines); row++ {
		curLine := lines[row]
		acc := 0
		pCnt := 0

		for i := 0; i < len(curLine); i++ {
			if curLine[i] == 'X' {
				acc = 0
			}
			if curLine[i] == '.' {
				acc += 1
			}
			if acc == 2 {
				pCnt += 1
			}
		}
		cnt += pCnt
	}

	return cnt
}

func reverseLines(lines []string, N int) []string {

	copied := make([]string, len(lines))
	copy(copied, lines)

	for col := 0; col < N; col++ {
		newLine := ""
		for row := 0; row < N; row++ {
			newLine += string(lines[row][col])
		}
		copied[col] = newLine
	}
	return copied
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
