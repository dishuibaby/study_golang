package main

import "fmt"

type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {
	//return strconv.Itoa(t.a) + "/" + strconv.FL
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}

type Celsius float64

func (c Celsius) String() string {
	//return strconv.Itoa(t.a) + "/" + strconv.FL
	return fmt.Sprintf("%.2f°C", c)
}

type Day int

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Day) String() string {
	if d < MO || d > SU {
		return "Unknown"
	}
	return days[d]
}

func main() {
	t := &T{7, -2.53, "abc\tdef"}
	fmt.Println("Ret", t)

	temp := Celsius(36.6)
	fmt.Println(temp)

	today := Day(FR)
	fmt.Println(today)
}
