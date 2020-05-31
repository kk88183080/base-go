package main

import "fmt"

func main() {
	//_24testok()
	_24TestInter()
}

// 于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，
// 由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错
/*func _24test()  {
	var i int =1
	i.printint()
}

func (i int) printint()  {
	fmt.Println(i)
}*/

type _24Myint int

func (i _24Myint) _24printint() {
	fmt.Println(i)
}

func _24testok() {
	var i _24Myint = 10
	i._24printint()
}

//
type _24people interface {
	Speak(string) string
}

type _24Student struct {
}

func (s *_24Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func _24TestInter() {
	var peo _24people = &_24Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
