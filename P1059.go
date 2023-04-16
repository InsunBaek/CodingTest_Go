// BruteForce
// https://www.acmicpc.net/problem/1059

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

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	L := scanInt()
	Set := make([]int, 0)
	for i := 0; i < L; i++ {
		Set = append(Set, scanInt())
	}
	n := scanInt()

	sort.Slice(Set, func(i, j int) bool {
		return Set[i] < Set[j]
	})

	sA, eB, flag := 0, 0, 0
	for i := 0; i < L - 1; i++ {
		if Set[i] < n && n < Set[i + 1] {
			sA, eB = Set[i], Set[i + 1]
			flag = 1
		}
	}

	if n < Set[0] {
		flag = 2
	}

	if flag == 0 {
		fmt.Println("0")
	} else if flag == 1 {
		fmt.Println(BruteForce(sA, eB, n))
	} else {
		// 좋은 구간 정의에 따라서 n이 Set[0] 보다 작은 경우의 수도 고려해야한다.
		fmt.Println(BruteForce(0, Set[0], n))
	}
}

func BruteForce(sA, eB, n int) int {
	minA := sA + 1
	maxB := eB - 1
	cnt := 0
	for curA := minA; curA <= n; curA++ {
		for curB := curA + 1; curB <= maxB; curB++ {
			if curA <= n && n < curB {
				cnt++
			} else if curA < n && n <= curB {
				cnt++
			}
		}
	}
	return cnt
} 


func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
