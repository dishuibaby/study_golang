package main

import (
	"fmt"
	"reflect"
)

func main() {
	fv := func() { fmt.Println("Hello World") }
	fv()
	fmt.Println("fv Type:", reflect.TypeOf(fv))

	fmt.Println(f())
}

func f() (ret int) {
	defer func() {
		fmt.Println(ret)
		ret++
	}()

	return 2
}
