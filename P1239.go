// BruteForce
// https://www.acmicpc.net/problem/1239

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int
var max int

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N = scanInt()
	max = 0
	src := make([]int, 0)

	for i := 0; i < N; i++ {
		src = append(src, scanInt())
	}

	sort.Slice(src, func(i, j int) bool {
		return src[i] > src[j]
	})

	// 1개라도 50초과이면 이등분선 불가능
	if src[0] > 50 {
		fmt.Fprintln(bw, "0")
	} else {
		for i := 0; i < 1; i++ {
			marking := make([]bool, 100)
			marking[0] = true
			used := make([]bool, N)
			used[i] = true

			DFS(i, 0, 0, 0, marking, used, src)
		}
		fmt.Println(max)
	}
}

func DFS(currentIndex, filled, depth, cnt int, marking, used []bool, src []int) {

	// 현재 뽑아낸 값
	curValue := src[currentIndex]
	// 현재 값을 더한 새로운 위치
	filled += curValue

	// 모든 Value 사용한 경우 return
	if depth == N - 1{
		if max < cnt {
			max = cnt
		}
		return
	}

	// 이등분선 유효성 검사 가능
	if filled >= 50 {
		// 이등분선 possible
		if marking[filled - 50] {
			cnt += 1
		}
	}

	// 마킹
	marking[filled] = true 
	used[currentIndex] = true 
	
	// 사용하지 않은 Value들 DFS
	for newIndex := 0; newIndex < N; newIndex++ {
		if used[newIndex] { // 사용한 값이라면 pass
			continue
		}
		DFS(newIndex, filled, depth + 1, cnt, marking, used, src)
	}

	// 마킹 해제
	marking[filled] = false
	used[currentIndex] = false
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
