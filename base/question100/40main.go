package main

import "fmt"

// 都正确
//A. select机制用来处理异步IO问题；
//B. select机制最大的一条限制就是每个case语句里必须是一个IO操作；
//C. golang在语言级别支持select关键字；

func main() {
	_40test2()
}

// 编译出错，有方向的chan 不能关闭
func _40test(i <-chan int) {
	//close(i)
}

type Param map[string]interface{}

type Show struct {
	*Param
}

//存在两个问题：1.map 需要初始化才能使用；2.指针不支持索引
func _40test1() {
	//s:= new(Show)
	//s.Param["l"] =2
}

// 问题修复
func _40test2() {
	s := new(Show)

	p := make(Param)
	p["day"] = 2
	s.Param = &p

	t := *s.Param
	fmt.Println(t["day"])
}
