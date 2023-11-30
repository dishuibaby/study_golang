package main

import "fmt"

func main() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			//fmt.Printf("before: i=%d, j=%d\n", i, j)
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("after: i=%d, j=%d\n", i, j)
		}
	}
}
