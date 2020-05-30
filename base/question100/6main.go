package main

import "fmt"

/*type Myint1 int
type Myint2 = int


func main() {
	var i int = 0
	var i1 Myint1 = Myint1(i)
	var i2 Myint2 = i
	// 下面这行代码编译出错
	//var i3 Myint1 = i
	fmt.Println(i1, i2)
}*/

//第 5 行代码是基于类型 int 创建了新类型 MyInt1，第 6 行代码是创建了 int 的类型别名 MyInt2，注意类型别名的定义时 = 。所以，
//第 10 行代码相当于是将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过；
//而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值。
//第 10 行代码的赋值可以使用强制类型转化 var i1 MyInt1 = MyInt1(i)

type user struct {
	id int
}

func main() {
	u := &user{
		id: 1,
	}

	fmt.Printf("%T", u)

	fmt.Printf("ud:%d\n", u.id)
	fmt.Println((*u).id)
}
