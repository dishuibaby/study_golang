package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example 中国 of a string"

	fmt.Printf("string \"%s\" have prefix %s?", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	fmt.Printf("check contain:%t\n", strings.Contains(str, "exa"))

	fmt.Printf("check Index: %d\n", strings.Index(str, "s"))
	fmt.Printf("check Last Index: %d\n", strings.LastIndex(str, "s"))
	fmt.Printf("check Last Index: %d\n", strings.IndexRune(str, rune('国')))
	fmt.Printf("check Last Index: %d\n", strings.Index(str, "国"))

	newStr := strings.Replace(str, "s", "*", -1)
	fmt.Println(newStr)

	unitStr := "Oh, "
	fmt.Println("count times:", strings.Count(str, "a"))
	newStr = strings.Repeat(unitStr, 3)
	fmt.Println(newStr)

	newStr = "This Is String"
	fmt.Println("To Lower:", strings.ToLower(newStr))
	fmt.Println("To Upper:", strings.ToUpper(newStr))

	newStr = " This is String "
	newStr = strings.TrimSpace(newStr)
	fmt.Println("Trim Space:", newStr)
	fmt.Println("Trim:", strings.Trim(newStr, "ing"))

	newStr = "Tian Ming Is One Xiaobing"
	newSlice := strings.Split(newStr, " ")
	fmt.Println("split", newSlice)
	fmt.Println("Join", strings.Join(newSlice, "-"))

}
