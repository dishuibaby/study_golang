package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Println("str1 Len:", len(str1))
	fmt.Println("str1 Rune Len:", utf8.RuneCountInString(str1))

	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Println("str2 Len:", len(str2))
	fmt.Println("str2 Rune Len:", utf8.RuneCountInString(str2))
}
