// BruteForce
// https://www.acmicpc.net/problem/1251

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn()
}

func Fn() {
	str := scanString()
	newStrSlice := make([]string, 0)

	for i := 1; i < len(str) - 1; i++ {
		for j := i + 1; j < len(str); j++ {
			newStr := ""
			newStr += reverse(0, i, str)
			newStr += reverse(i, j, str)
			newStr += reverse(j, len(str), str)

			newStrSlice = append(newStrSlice, newStr)
		}
	}
	
	sort.Slice(newStrSlice, func(i, j int) bool {
		return (newStrSlice[i] < newStrSlice[j])
	})
	fmt.Println(newStrSlice[0])
}

func reverse(start, end int, inStr string) string {
	s := ""
	for i := end - 1; i >= start; i-- {
		s += string(inStr[i])
	}
	return s
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}