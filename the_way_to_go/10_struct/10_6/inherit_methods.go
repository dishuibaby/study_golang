package main

import "fmt"

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(newId int) {
	b.id = newId
}

type Person struct {
	Base
	FisrtName string
	LastName  string
}

type Employee struct {
	Person
	Salary int
}

func main() {
	person := Person{}
	person.Base = Base{}
	person.SetId(1)
	person.FisrtName = "Tian"
	person.LastName = "Ming"
	employee := Employee{person, 1000}

	fmt.Println(employee.Id())
}
