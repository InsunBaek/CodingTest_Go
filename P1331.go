// Simulation
// https://www.acmicpc.net/problem/1331

package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn()
}

func Fn() {
	board := make([][]bool, 6)
	flag := true
	
	for i := 0; i < 6; i++ {
		board[i] = make([]bool, 6)
	}

	start_X, start_Y := scanSpecial()
	board[start_Y][start_X] = true
	current_X, current_Y := start_X, start_Y

	for i := 0; i < 35; i++ {
		new_X, new_Y := scanSpecial()
		// 이미 답이 나왔으면 continue
		if !flag {
			continue
		}

		// 현재 좌표에서 새로운 좌표를 나이트 점프로 갈수 없다면 flag 수정하고 continue
		if canNotReach(current_X, current_Y, new_X, new_Y) {
			flag = false
			continue
		}

		// 방문하려는 좌표가 이미 방문한 곳이라면
		if board[new_Y][new_X] {
			flag = false
			continue
		}

		// 모든 조건을 통과 했다면 방문처리를 해주고 current 좌표 업데이트
		board[new_Y][new_X] = true
		current_X, current_Y = new_X, new_Y
	}

	
	if flag {
		// flag가 살아있다면 current->start 도달 여부 판단으로 최종 결정
		if canNotReach(current_X, current_Y, start_X, start_Y) {
			fmt.Println("Invalid")
		} else {
			fmt.Println("Valid")
		}
	} else {
		fmt.Println("Invalid")
	}

}

func canNotReach(curX, curY, newX, newY int) bool {
	distance_X := Abs(curX - newX)
	distance_Y := Abs(curY - newY)

	if distance_X + distance_Y != 3 {
		return true
	} else {
		if distance_X == 0 || distance_Y == 0 {
			return true
		} else {
			return false
		}
	}
}

func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func scanSpecial() (int, int) {
	sc.Scan()
	X, Y := sc.Text()[0], sc.Text()[1]
	return int(X - 65), int(Y - 49)
}