package main

import (
	"fmt"

	"./sort"
)

func ints() {
	data := []int{2134, 4345, 62, 31, 4, 54, 3, 555, 23, 24, 23425, 6, 46}
	a := sort.IntArr(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

func strings() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	a := sort.StringArr(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
	_bool := sort.StringsAreSorted(a)
	fmt.Println(_bool)
}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArr struct {
	data []*day
}

func (p *dayArr) Len() int           { return len(p.data) }
func (p *dayArr) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArr) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}

	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArr{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("fail")
	}

	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}

func main() {
	ints()
	strings()
	days()
}
