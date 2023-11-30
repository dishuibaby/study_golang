package main

import "fmt"

func main() {
	res := 0
	for i := 0; i <= 10; i++ {
		res = fibonacci(i)
		fmt.Println(res)
	}

	fibFunc := fibonacci2()
	for i := 0; i <= 10; i++ {
		fmt.Println(fibFunc())
	}
}

func fibonacci2() func() int {
	a, b := 0, 1
	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	return res
}
