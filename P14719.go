// Simulation
// https://www.acmicpc.net/problem/14719

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
	_, M := scanInt(), scanInt()
	blocks := make([]int, 0)
	
	for i := 0; i < M; i++ {
		blocks = append(blocks, scanInt())
	}
	
	answer := 0
	
	for i := 1; i < M-1; i++ {
		left := 0
		right := 0
	
		// 왼쪽 최대 값
		for j := 0; j < i; j++ {
			left = max(left, blocks[j])
		}
	
		// 오른쪽 최대 값
		for j := M - 1; j > i; j-- {
			right = max(right, blocks[j])
		}
	
		// 빗물 계산
		answer += max(0, min(left, right)-blocks[i])
	}
	
	fmt.Println(answer)
}
	
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
	
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
