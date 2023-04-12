// Stack
// https://www.acmicpc.net/submit/1662

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	stack := *NewStack()
	str := scanString()

	
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" {
			// "("는 무시
			stack.Push(-1)
		} else if string(str[i]) == ")" {
			cnt := 0;
			for {
				temp := stack.Pop().(int)
				if temp == -1 {
					break
				}
				cnt += temp
			}
			stack.Push(stack.Pop().(int) * cnt)
		} else if i < len(str) - 1 && string(str[i + 1]) == "(" {
			// "(" 앞 계수로 들어가야되는 숫자
			num, _ := strconv.Atoi(string(str[i]))
			stack.Push(num)
		} else {
			// 일반 연속적인 숫자일때
			stack.Push(1);
		}
	}

	answer := 0;
	for !stack.Empty() {
		answer += stack.Pop().(int)
	}
	fmt.Println(answer)
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

// Stack

type Stack struct {
	stack *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (st *Stack) Push(item interface{}) {
	st.stack.PushBack(item)
}

func (st *Stack) Pop() interface{} {
	top := st.stack.Back()
	if top == nil {
		return nil
	}
	return st.stack.Remove(top)
}

func (st *Stack) Empty() bool {
	if st.stack.Len() == 0 {
		return true
	} else {
		return false
	}
}