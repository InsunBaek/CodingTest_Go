// Greedy
// https://www.acmicpc.net/problem/1080

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var n, m int
var t int = 0
var board, ans [][]int

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	n, m = scanInt(), scanInt()
	
	board = make([][]int, n)
	ans = make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		ans[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		line := scanString()
		for j := 0; j < m; j++ {
			token, _ := strconv.Atoi(string(line[j]))
			board[i][j] = token
		}
	}
	for i := 0; i < n; i++ {
		line := scanString()
		for j := 0; j < m; j++ {
			token, _ := strconv.Atoi(string(line[j]))
			ans[i][j] = token
		}
	}

	for i := 0; i < n - 2; i++ {
		for j := 0; j < m - 2; j++ {
			if board[i][j] != ans[i][j] {
				t += 1
				flip(i, j)
			}
		}
	}

	flag := false;
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != ans[i][j] {
				flag = true
				break
			}
		}
	}
	if flag {
		fmt.Println("-1")
	} else {
		fmt.Println(t)
	}
}

func flip(x, y int) {
	if x+3>n || y+3>m {
		return
	}
	for i := x; i < x + 3; i++ {
		for j := y; j < y + 3; j++ {
			if board[i][j] == 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
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
