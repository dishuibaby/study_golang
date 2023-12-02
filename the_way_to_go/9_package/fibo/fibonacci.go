// Package fibo fibo/fibonacci.go
package fibo

var lastResult int

func Fibonacci(n int, op string) int {
	if n <= 1 {
		lastResult = 1
		return 1
	}

	switch op {
	case "+":
		lastResult = Fibonacci(n-1, op) + Fibonacci(n-2, op)
	case "*":
		lastResult = Fibonacci(n-1, op) * Fibonacci(n-2, op)
	default:
		lastResult = 0
	}

	return lastResult
}

func GetLastResult() int {
	return lastResult
}
