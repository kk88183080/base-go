package main

import "fmt"

func main() {
	t := _13Teacher{}
	t.showA()
}

type _13Person struct {
	id int
}

func (p *_13Person) showA() {
	fmt.Println("showA")
	p.showB()
}

func (*_13Person) showB() {
	fmt.Println("showB")
}

type _13Teacher struct {
	_13Person
}

func (*_13Teacher) showB() {
	fmt.Println("teacher showB")
}
