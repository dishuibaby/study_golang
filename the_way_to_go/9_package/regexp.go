package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//目标字符串
	searchIn := "小六: 2578.34 小糖: 4567.23 小饼: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	if ok, _ := regexp.MatchString(pat, searchIn); ok {
		fmt.Println("Match Succ.")
	}

	re, _ := regexp.Compile(pat)

	//替换匹配到的内容
	str := re.ReplaceAllString(searchIn, "**.*")
	fmt.Println(str)

	//参数为函数
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
