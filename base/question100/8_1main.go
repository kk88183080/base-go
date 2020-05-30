package main

import "fmt"

func main() {
	// 基本类型 不可以使用.(type), 结构体也不行，只在接口类型才可以
	//i := GetValue()
	// 这个是返回的指针的结构体
	//i := GetUser()
	//i := GetPerson()

	// 只有这一个编译可以通过
	i := GetBoy()
	switch i.(type) {

	case int:
		fmt.Println("int")
		break
	case string:
		fmt.Println("string")
		break
	case interface{}:
		fmt.Println("interface{}")
		break
	default:
		fmt.Println("default")
		break
	}
}

type user1 struct {
	id int
}

type boy struct {
	base user1
}

type person interface {
	hello() string
}

func GetValue() int {
	return 1
}

func GetUser() *user1 {
	return &user1{id: 10}
}

func (*boy) hello() string {
	return "boy hello"
}

func GetPerson() *boy {
	b := &boy{}
	b.base.id = 10
	return b
}

func GetBoy() interface{} {
	b := &boy{}
	b.base.id = 10
	return b
}
