package main

import "fmt"

func main() {
	_19testdefer() //
}

type _19person struct {
	age int
}

func _19testdefer() {

	person := &_19person{1}

	defer fmt.Println(person.age)

	defer func(p *_19person) {
		fmt.Println(p.age)
	}(person)

	defer func() {
		fmt.Println(person.age)
	}()

	// 30 30 1
	//person.age = 30

	// 40 1 1
	person = &_19person{40}
}
