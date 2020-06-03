package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//_49test()
	_49test2()
}

func _49test() {
	var ch chan int
	select {
	case v, ok := <-ch:
		fmt.Println(v, ok)
	default:
		fmt.Printf("default")

	}
}

type _49People struct {
	//结构体访问控制，因为 name 首字母是小写，导致其他包不能访问，所以输出为空结构体
	//name string `json:"name"`
	Name string `json:"name"`
}

func _49test2() {

	js := `{
	"name":"ljw"	
	}`

	var p _49People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil { //失败

		fmt.Println(err)
		return
	}

	fmt.Println(p)
}
