package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice1()

	arr := [5]int{1, 2, 3, 4, 5}
	arrSum := sum(arr[:])
	fmt.Println(arrSum)

	fmt.Println("-------------------")
	makeSlice()
	fmt.Println("-------------------")

	s1 := []byte{1, 2, 3, 4}
	data := []byte{5, 6, 7, 8}
	result := Append(s1, data)
	fmt.Println(result)
	fmt.Println("-------------------")
	SplitBuf()
	fmt.Println("-------------------")
	//s := make([]int, 10, 20)
	//s := []int{122, 234, 35, 12, 213, 4121, 522}
	//
	//min := MinSlice(s)
	//fmt.Println(min)
	//fmt.Println("-------------------")

	s := make([]int, 10, 20)
	rand.Seed(time.Now().UnixNano())
	for i := range s {
		s[i] = (i + 1) * rand.Intn(100)
	}
	fmt.Println(s)
	fmt.Println("-------------------")
	reslicing()
	fmt.Println("-------------------")
	copySlice()
}

func slice1() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(string(b[1:4])) //ola
	fmt.Println(string(b[:2]))  //ol
	fmt.Println(string(b[2:]))  //
	fmt.Println(string(b[:]))
}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}

	return s
}

func makeSlice() {
	//var s1 []int = make([]int, 10)
	s1 := make([]int, 10)
	for i := range s1 {
		s1[i] = 5 * i
	}

	for i, v := range s1 {
		fmt.Printf("Slice at %d is %d\n", i, v)
	}

	fmt.Println("Slice Len:", len(s1))
	fmt.Println("Slice Cap:", cap(s1))
}

func Append(slice, data []byte) []byte {
	return append(slice, data...)
}

func SplitBuf() {
	buf := []byte{1, 2, 3, 4, 5, 6}
	n := 4
	first, second := buf[:n], buf[n:]

	fmt.Println(first)
	fmt.Println(second)
}

func MinSlice(s []int) int {
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

func reslicing() {
	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

func copySlice() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
