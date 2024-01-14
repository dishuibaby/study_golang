package main

import "fmt"

type Simpler interface {
	Get() string
	Set(name string)
}

type Simple struct {
	name string
}

func (s Simple) Get() string {
	return s.name
}

func (s *Simple) Set(name string) {
	s.name = name
}

func Demo(s Simpler) {
	s.Set("小饼")
	name := s.Get()
	fmt.Println(name)
}

func main() {
	var s Simpler = &Simple{}
	Demo(s)
}
