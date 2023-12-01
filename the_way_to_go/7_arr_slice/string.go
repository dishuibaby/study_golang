package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

func main() {
	strLen()
	fmt.Println("-------------------")

	changeStr()
	fmt.Println("-------------------")

	a := "asdcascdsfasdfdfs"
	b := "aaaaaaaaaaaaaaaaaaaaaaaaa"

	ret := compare([]byte(a), []byte(b))
	fmt.Println(ret)
	fmt.Println("-------------------")

	sorting()
}

func strLen() {
	str1 := "Study Golang!"
	str2 := "我正在学习Golang"

	fmt.Println("str1 len of byte:", len(str1))
	fmt.Println("str1 len of char:", utf8.RuneCountInString(str1))
	fmt.Println("str2 len of byte:", len(str2))
	fmt.Println("str2 len of char:", utf8.RuneCountInString(str2))
}

func changeStr() {
	s := "你好 Go"
	c := []rune(s)
	c[0] = '您'
	s2 := string(c)
	fmt.Println(s2)
}

func compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}

	return 0
}

func sorting() {
	slice1 := []int{12, 2, 3, 134, 1, 325, 632, 12, 33}
	sort.Ints(slice1)
	fmt.Println(slice1)
	fmt.Println(sort.IntsAreSorted(slice1))
}

func practice() {
	//7.12 编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。

}
