package main

import "fmt"

func main() {

	//test11()
	test11_2()
}

func test11_2() {
	s := make(map[string]int)

	delete(s, "h")

	fmt.Println(s["h"])
}

//当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。
func test11() {
	var i interface{}
	if i == nil {
		fmt.Println("is nil")
		return
	}

	fmt.Println("not is nil")
}
