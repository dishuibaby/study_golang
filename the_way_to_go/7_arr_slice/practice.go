package main

import (
	"fmt"
	"strings"
)

func main() {
	//7.12
	str := "Golang arr slice string 练习"
	str1, str2 := practice712([]rune(str), 11)
	fmt.Println(string(str1), "---", string(str2))

	//7.13
	fmt.Println("-=-=-=-=-=")
	practice713()

	fmt.Println("-=-=-=-=-=")
	practice714()

	fmt.Println("-=-=-=-=-=")
	practice715()

	fmt.Println("-=-=-=-=-=")
	slice := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original slice:", slice)
	practice716(slice)
	fmt.Println("Sorted slice:", slice)

	fmt.Println("-=-=-=-=-=")
	nums := []int{1, 3, 5, 7, 9}
	multiplied := mapFunc(multiplyByTen, nums)
	fmt.Println(multiplied)
}

// 编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
func practice712(str []rune, i int) (str1, str2 []rune) {
	if i > len(str) {
		i = len(str)
	}

	str1, str2 = str[:i], str[i:]

	return
}

// 假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？
func practice713() {
	str := "Golang Study 练习中。"
	str2 := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(str2)

	runes := []rune(str)
	half := len(runes) / 2

	str3 := string(runes[half:]) + string(runes[:half])
	fmt.Println(str3)
}

/*
* 编写一个程序，要求能够反转字符串，即将 “Google” 转换成 “elgooG”（提示：使用 []byte 类型的切片）。
* 如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）。
* 如果您想要反转 Unicode 编码的字符串，请使用 []int32 类型的切片。
 */
func practice714() {
	str := "Golang 学习"
	runes := []rune(str)

	// 原地反转
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 将 rune 切片转换回字符串
	newStr := string(runes)

	fmt.Println(newStr)        // 输出反转后的字符串
	fmt.Println(string(runes)) // 输出反转后的字符串

	var builder strings.Builder
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}

	newStr = builder.String()
	fmt.Println(newStr) // 输出反转后的字符串
}

// 编写一个程序，要求能够遍历一个字符数组，并将当前字符和前一个字符相同的字符拷贝至另一个数组。
func practice715() {
	str := []rune("this is book, u can seek it")
	repeatStr := make([]rune, 10, 50)

	for i, v := range str {
		if i == 0 {
			continue
		}
		if str[i-1] == v {
			repeatStr = append(repeatStr, v)
		}
	}

	fmt.Println(string(repeatStr))
}

// 编写一个程序，使用冒泡排序的方法排序一个包含整数的切片
func practice716(slice []int) {
	len := len(slice)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

type IntToIntFunc func(int) int

/*
*
在函数式编程语言中，一个 map-function 是指能够接受一个函数原型和一个列表，
并使用列表中的值依次执行函数原型，
公式为：map ( F(), (e1,e2, . . . ,en) ) = ( F(e1), F(e2), ... F(en) )。

编写一个函数 mapFunc 要求接受以下 2 个参数：
一个将整数乘以 10 的函数
一个整数列表

最后返回保存运行结果的整数列表。
*/
func mapFunc(f IntToIntFunc, list []int) []int {
	result := make([]int, 0)
	for _, v := range list {
		result = append(result, f(v))
	}

	return result
}

func multiplyByTen(n int) int {
	return n * 10
}
