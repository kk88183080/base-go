package main

import (
	"./ArrayList"
	"./StackArray"
	"fmt"
)

func main() {

	//test4()
	//test5()
	//testStack()
	//testArrayStack()
	//testArrayStackX()
	//testArrayStackXIterator()

	//testSumNum()
	//testSumStack()
	//testFeb()
	//testFebStack()
	//testFebStackDesc()
	//testFiles()
	//testFilesStack()

	//testFilesTree()
	testFilesDeep()
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

func testStack() {

	s := StackArray.NewStack()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	fmt.Println(s)

	for i := 0; i < 10; i++ {
		fmt.Println(s.Pop())
	}
	fmt.Println(s)
}

func testArrayStack() {

	s := ArrayList.NewArrayListStack()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	fmt.Println(s)

	for i := 0; i < 10; i++ {
		fmt.Println(s.Pop())
	}
	fmt.Println(s)
}

func testArrayStackX() {

	s := ArrayList.NewArrayListStackX()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	fmt.Println(s)

	for i := 0; i < 10; i++ {
		fmt.Println(s.Pop())
	}
	fmt.Println(s)
}

func testArrayStackXIterator() {

	s := ArrayList.NewArrayListStackX()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	fmt.Println(s)

	for it := s.Myit; it.HashNext(); {
		next, _ := it.Next()
		fmt.Println(next)
	}
	fmt.Println(s)
}

func testSumNum() { // 5+4+3 + 2 + 1
	fmt.Println(StackArray.Sum(5))
}

func testSumStack() { // 5=15 3=6
	fmt.Println(StackArray.SumStack(6))
}

func testFeb() {
	fmt.Println(StackArray.Ferbi(9))
}

func testFebStack() {
	fmt.Println(StackArray.FerbiStack(6))
}

func testFebStackDesc() {
	fmt.Println(StackArray.FerbiStackDesc(9))
}

/**
查询全部的文件
*/
func testFiles() {

	path := "/Users/liangweili/Desktop/bigdata-workspace/base-go/base/DataStructOption"
	files := []string{}
	all, _ := StackArray.GetAll(path, files)

	for i := 0; i < len(all); i++ {
		fmt.Println(i, ">", all[i])
	}
}

func testFilesStack() {
	fmt.Println("stack impl")

	path := "/Users/liangweili/Desktop/bigdata-workspace/base-go/base/DataStructOption"
	stack := StackArray.GetAllStack(path)

	for i, v := range stack {
		fmt.Println(i, ">", v)
	}
}

/**
查询全部的文件
*/
func testFilesTree() {

	path := "/Users/liangweili/Desktop/bigdata-workspace/base-go/base/DataStructOption"
	files := []string{}
	all, _ := StackArray.GetAllX(path, files, 1)

	for i := 0; i < len(all); i++ {
		fmt.Println(all[i])
	}
}

func testFilesDeep() {
	path := "/Users/liangweili/Desktop/bigdata-workspace/base-go/base/DataStructOption"
	StackArray.QueryFileDeep(path, 1)
}
