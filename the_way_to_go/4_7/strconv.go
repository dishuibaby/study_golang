package main

import (
	"fmt"
)

func main() {
	//i := 123
	//fmt.Printf("i str: %T\n", strconv.Itoa(i))
	//
	//fmt.Println(time.Now().Unix())
	//fmt.Println(time.Now().UnixMicro())
	//
	//now := time.Now()                                  // 获取当前时间
	//millis := now.UnixNano() / int64(time.Millisecond) // 转换为毫秒时间戳
	//fmt.Println(time.Millisecond)
	//fmt.Println(int64(time.Millisecond))
	//fmt.Println("Milliseconds Timestamp:", millis)

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
