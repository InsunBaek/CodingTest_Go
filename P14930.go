// Sort
// https://www.acmicpc.net/problem/14930

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

type BEAD struct {
	pos int
	vel int
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, T := scanInt(), scanInt()
	BEADs := make([]BEAD, 0)

	for i := 0; i < N; i++ {
		BEADs = append(BEADs, BEAD{pos: scanInt(), vel: scanInt()})
	}

	redBead := BEADs[0]
	sort.Slice(BEADs, func(i, j int) bool {
		return BEADs[i].pos < BEADs[j].pos
	})
	redIdx := -1
	for i := 0; i < len(BEADs); i++ {
		if BEADs[i] == redBead {
			redIdx = i
		}
	}

	vPos := make([]int, N)
	for i := 0; i < N; i++ {
		vPos[i] = BEADs[i].pos + BEADs[i].vel * T
	}
	sort.Slice(vPos, func(i, j int) bool {
		return vPos[i] < vPos[j]
	})

	fmt.Println(vPos[redIdx])
}

func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
