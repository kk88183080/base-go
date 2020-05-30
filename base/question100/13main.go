package main

import "fmt"

func main() {
	//_13test()

	_13test1()
}

func _13test() {
	i := 5
	defer _13hello(i)
	i = i + 10
}

func _13test1() {
	i := 5
	defer _13hello(i)
	i += 10
	defer _13hello(i)

}

func _13hello(i int) {
	fmt.Println(i)
}

//这个例子中，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5
