package main

import (
	"fmt"
	"study_golang/the_way_to_go/10_6/person"
)

func main() {
	p := new(person.Person)
	p.SetFirstName("小饼")
	fmt.Println(p.FisrtName())
}
