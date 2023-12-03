package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	filed1 bool   "1st Tag"
	filed2 string "2nd Tag"
	filed3 int    "3nd Tag"
}

func main() {
	tagType := TagType{true, "小饼", 3}
	for i := 0; i < 3; i++ {
		refTag(tagType, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)

	fmt.Printf("%v\n", ixField.Tag)
}
