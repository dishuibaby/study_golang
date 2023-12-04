package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:\n" + fmt.Sprintln(c.Log.String())
}

func main() {
	c := &Customer{"小饼", Log{"1st Msg"}}
	c.Add("2nd Msg")
	fmt.Println(c)
}
