// Simulation
// https://www.acmicpc.net/problem/27918

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
	dPoint, pPoint := 0, 0
	inputs := make([]string, 0)
	for i := 0; i < N; i++ {
		inputs = append(inputs, scanString())
	}
	for i := 0; i < N; i++ {
		roundWinner := inputs[i]
		if roundWinner == "D" {
			dPoint++
		}
		if roundWinner == "P" {
			pPoint++
		}

		if dPoint - pPoint >= 2 || pPoint - dPoint >= 2{
			fmt.Print(dPoint)
			fmt.Print(":")
			fmt.Print(pPoint)
			break;
		}
		if i == N - 1 {
			fmt.Print(dPoint)
			fmt.Print(":")
			fmt.Print(pPoint)
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
