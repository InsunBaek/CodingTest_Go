// Prefix Sum
// https://www.acmicpc.net/problem/14929

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

	n := scanInt()
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, scanInt())
	}

	prefix_sum, res := 0, 0
	for i := 0; i < len(nums) - 1; i++ {
		prefix_sum += nums[i]
		res += nums[i + 1] * prefix_sum
	}

	fmt.Println(res)
}


func scanInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}
