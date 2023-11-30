package main

import "fmt"

func main() {
	f := Adder()            // x=0
	fmt.Print(f(1), " - ")  //echo 1, delta=1, 运行之后x=1
	fmt.Print(f(20), " - ") //echo 21, delta = 20, 运行之后x=21
	fmt.Print(f(300))       //echo 321, delta = 300

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
	}(1000) // Passes argument 1000 to the function literal.
	fmt.Println(g)

}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
