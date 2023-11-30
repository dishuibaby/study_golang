package main

import "fmt"

func main() {
	x := min(4, 5, 6, 34, 23)
	fmt.Println(x)

	slice := []int{33, 23, 13, 32, 554, 32}
	min := min(slice...)
	fmt.Println(min)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}

	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}
