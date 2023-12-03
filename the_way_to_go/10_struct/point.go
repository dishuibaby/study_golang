package main

import (
	"fmt"
	"math"
)

// Point 结构体表示二维点
type Point struct {
	X, Y float64
}

// Abs 方法计算点表示的向量的长度
func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Scale 方法将点的坐标乘以一个尺度因子
func (p *Point) Scale(scaleFactor float64) {
	p.X *= scaleFactor
	p.Y *= scaleFactor
}

// Polar 结构体表示三维点的极坐标
type Polar struct {
	R, Theta, Phi float64
}

func main() {
	p := Point{X: 3, Y: 4}
	fmt.Printf("Original Point: %+v\n", p)

	length := p.Abs()
	fmt.Printf("Length of the vector: %v\n", length)

	scaleFactor := 2.0
	p.Scale(scaleFactor)
	fmt.Printf("Scaled Point: %+v\n", p)
}
