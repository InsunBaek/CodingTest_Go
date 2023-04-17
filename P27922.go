// Greedy
// https://www.acmicpc.net/problem/27922

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

type Triple struct {
	a, b, c int
}
func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, K := scanInt(), scanInt()
	tripleSlice := make([]Triple, 0)
	for i := 0; i < N; i++ {
		tripleSlice = append(tripleSlice, Triple{a: scanInt(), b: scanInt(), c: scanInt()})
	}

	calcMaxValue(tripleSlice, N, K)
}

func calcMaxValue(tripleSlice []Triple, N, K int) {
	type Pair struct {
		label, value int
	}
	
	pairSlice1 := make([]Pair, 0)
	pairSlice2 := make([]Pair, 0)
	pairSlice3 := make([]Pair, 0)

	for i := 0; i < N; i++ {
		// 3개 과목중 2개 선택하여 정렬
		pairSlice1 = append(pairSlice1, Pair{label: i, value: tripleSlice[i].a + tripleSlice[i].b})
		pairSlice2 = append(pairSlice2, Pair{label: i, value: tripleSlice[i].a + tripleSlice[i].c})
		pairSlice3 = append(pairSlice3, Pair{label: i, value: tripleSlice[i].b + tripleSlice[i].c})
	}

	sort.Slice(pairSlice1, func(i, j int) bool {
		return pairSlice1[i].value > pairSlice1[j].value
	})
	sort.Slice(pairSlice2, func(i, j int) bool {
		return pairSlice2[i].value > pairSlice2[j].value
	})
	sort.Slice(pairSlice3, func(i, j int) bool {
		return pairSlice3[i].value > pairSlice3[j].value
	})
	
	visited1 := make([]bool, N)
	visited2 := make([]bool, N)
	visited3 := make([]bool, N)

	sum1, sum2, sum3 := 0, 0, 0

	for i := 0; i < K; i++ {
		curLabel1 := pairSlice1[i].label
		if visited1[curLabel1] {
			continue
		}
		visited1[curLabel1] = true
		sum1 += pairSlice1[i].value
	}

	for i := 0; i < K; i++ {
		curLabel2 := pairSlice2[i].label
		if visited2[curLabel2] {
			continue
		}
		visited2[curLabel2] = true
		sum2 += pairSlice2[i].value
	}

	for i := 0; i < K; i++ {
		curLabel3 := pairSlice3[i].label
		if visited3[curLabel3] {
			continue
		}
		visited3[curLabel3] = true
		sum3 += pairSlice3[i].value
	}
	fmt.Println(Max(sum1, sum2, sum3))	
}

func Max(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
