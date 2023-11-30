package main

import "fmt"

func main() {
	arr := [3]float64{3.4, 54.2, 60}
	x := Sum(&arr)
	fmt.Printf("The Sum of the arr is :%f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}

	return
}
