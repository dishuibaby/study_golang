package main

import (
	"fmt"
	"strings"
)

const LIMIT = 4

type StackArr [LIMIT]int

func (s *StackArr) Push(v int) {
	for i, val := range s {
		if val == 0 {
			s[i] = v
			return
		}
	}
}

func (s *StackArr) Pop() int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 0 {
			v := s[i]
			s[i] = 0
			return v
		}
	}

	return 0
}

func (s StackArr) String() string {
	str := strings.Builder{}
	for i, v := range s {
		str.WriteString(fmt.Sprintf("[%d:%d]", i, v))
	}

	return str.String()
}

func main() {
	var stack StackArr
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println("Stack:", stack)

	pop := stack.Pop()
	fmt.Println("Pop:", pop)
	fmt.Println("Stack:", stack)
}
