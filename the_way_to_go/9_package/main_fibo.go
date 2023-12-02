// main_fibo.go
package main

import (
	"./fibo"
	"fmt"
)

func main() {
	fmt.Println("Fibonacci with addition:", fibo.Fibonacci(10, "+"))
	fmt.Println("Last result:", fibo.GetLastResult())

	fmt.Println("Fibonacci with multiplication:", fibo.Fibonacci(10, "*"))
	fmt.Println("Last result:", fibo.GetLastResult())
}
