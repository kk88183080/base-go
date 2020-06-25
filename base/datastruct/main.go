package main

import (
	"./ArrayList"
	"fmt"
)

func main() {

	//test4()
	test5()
}

func test1() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append("sss")
	fmt.Println(list)
}

func test2() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append("sss")
	fmt.Println(list)
}

func test3() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append("sss")
	list.Delete(0)
	fmt.Println(list)
}

func test4() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append("sss")
	list.Append("3")
	list.Append("ss4s")
	list.Append("ss5s")
	list.Append("ss6s")
	list.Append("s7ss")
	list.Append("s8ss")
	list.Insert(0, "465")
	list.Insert(2, "xx465")
	list.Insert(3, "dd465")
	for i := 0; i < 100; i++ {
		list.Insert(4, fmt.Sprintf("aa465-%d", i))
	}
	fmt.Println(list)
}

// 测试迭代器
func test5() {
	var list ArrayList.List = ArrayList.NewArrayList()

	for i := 0; i < 30; i++ {
		list.Append(i)
	}

	for iterator := list.Iterator(); iterator.HashNext(); {
		next, _ := iterator.Next()
		fmt.Printf("index:=%d,val:=%v\n", iterator.GetIndex(), next)
	}

	for iterator := list.Iterator(); iterator.HashNext(); {
		next, _ := iterator.Next()
		fmt.Printf("del index:=%d,val:=%v\n", iterator.GetIndex(), next)
		iterator.Remove()
	}

	fmt.Println("输出list")
	fmt.Println(list)
}
