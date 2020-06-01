package main

import "fmt"

func main() {

	//_38test()
	//_38test1()

	_38test3()

}

func _38test() {

	x := []string{"a", "b", "c"}

	for v := range x {
		fmt.Print(v)
	}
}

func _38test1() {
	x := []string{"a", "b", "c"}
	for _, v := range x {
		fmt.Print(v)
	}
}

type _38User struct {
}

type _38User1 _38User

type _38User2 = _38User

func (i _38User1) m1() {
	fmt.Println("m1")
}

func (i _38User) m2() {
	fmt.Println("m2")
}

func _38test3() {
	var i1 _38User1
	var i2 _38User2

	i1.m1()
	i2.m2()

	// 编译出错
	//i1.m2()
	//i2.m1()
}

//第 2 行代码基于类型 User 创建了新类型 User1，第 3 行代码是创建了 User 的类型别名 User2，注意使用 = 定义类型别名。
//因为 User2 是别名，完全等价于 User，所以 User2 具有 User 所有的方法。
//但是 i1.m2() 是不能执行的，因为 User1 没有定义该方法。
