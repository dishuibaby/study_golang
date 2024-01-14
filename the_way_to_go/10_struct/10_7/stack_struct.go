package main

import (
	"fmt"
	"strings"
)

const LIMIT = 4

type StackStruct struct {
	index int
	data  [LIMIT]int
}

func (s *StackStruct) Push(v int) {
	if s.index >= LIMIT {
		return
	}

	s.data[s.index] = v
	s.index++
}

func (s *StackStruct) Pop() int {
	if s.index <= 0 {
		return 0
	}
	s.index--
	v := s.data[s.index]
	s.data[s.index] = 0
	return v
}

func (s StackStruct) String() string {
	var str strings.Builder
	for i, v := range s.data {
		str.WriteString(fmt.Sprintf("[%d:%d] ", i, v))
	}
	return str.String()
}

func main() {
	var stack StackStruct
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Stack:", stack)

	pop := stack.Pop()
	fmt.Println("Pop:", pop)
	fmt.Println("Stack:", stack)
}
