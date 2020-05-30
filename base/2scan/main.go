package main

import "fmt"

func main() {
	//Scan()
	//Scanln()
	Scanf()
}

/**
接受控制台输入的参数,把回车当空白处理
*/
func Scan() {
	a, b, c := "", 0, false
	fmt.Scan(&a, &b, &c)

	fmt.Println(a, b, c)
}

/**
输入回车，停止扫描
*/
func Scanln() {
	a, b, c := 1, "", true

	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)
}

func Scanf() {
	a, b, c := "", 0, false
	fmt.Scanf("%s, %s, %t", &a, &b, &c)
	fmt.Println(a, b, c)
}
