package main

import "fmt"

type Rectangle struct {
	W, H int
}

// 计算面积
func (r *Rectangle) Area() int {
	area := r.W * r.H
	return area
}

// 计算周长
func (r *Rectangle) Perimeter() int {
	perimeter := (r.W + r.H) * 2
	return perimeter
}

func main() {
	r := &Rectangle{300, 400}

	//面积
	area := r.Area()
	fmt.Println(area)

	//周长
	perimeter := r.Perimeter()
	fmt.Println(perimeter)
}
