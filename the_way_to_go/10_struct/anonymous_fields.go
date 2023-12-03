package main

import "fmt"

type innerS struct {
	int1 int
	int2 int
}

type outerS struct {
	a int
	b string
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.a = 12
	outer.b = "中国"
	outer.int = 99
	outer.int1 = 20
	outer.int2 = 30

	fmt.Printf("outer.a is: %d\n", outer.a)
	fmt.Printf("outer.b is: %s\n", outer.b)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.int1)
	fmt.Printf("outer.in2 is: %d\n", outer.int1)

	fmt.Println(outer)
}
