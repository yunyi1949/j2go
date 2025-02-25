package main

import "fmt"

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type Printer interface {
	Print()
}

func (m manager) Print() {
	fmt.Println("%+v\n", m)
}
func main5() {
	var m manager
	m.title = "java"
	m.name = "fwq"
	m.age = 30
	fmt.Println(m)
	fmt.Println(m.ToString())

	var p Printer = m

	p.Print()
}
