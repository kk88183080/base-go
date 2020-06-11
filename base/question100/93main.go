package main

import "fmt"

func main() {

}

type _93User struct {
	Name string
}

func (u _93User) SetName(name string) {

	u.Name = name
	fmt.Println(u.Name)
}

type Employee _93User

/*func _93test()  {
	emp := new(Employee)
	emp.SetName("l")
}*/

// 编译不通过。当使用 type 声明一个新类型，它不会继承原有类型的方法集。
