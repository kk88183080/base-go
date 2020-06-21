package main

import (
	"errors"
	"fmt"
)

func main() {

	s := &Stack{MaxTop: 5, top: -1}
	for i := 0; i < s.MaxTop; i++ {
		s.push(i + 1)
		fmt.Printf("push %d data:%v\n", i, s.data)
	}

	for i := 0; i < s.MaxTop; i++ {
		v, e := s.pop()
		if e != nil {
			fmt.Println("出栈出错了：" + e.Error())
		} else {
			fmt.Printf("%d, v=%d\n", i, v)
			fmt.Printf("%d, data=%v\n", i, s.data)
		}
	}

}

type Stack struct {
	MaxTop int    // 最多可以放多少个元素
	top    int    // 栈顶的下标，-1时表示栈为空
	data   [5]int // 存放数据的数组
}

func (s *Stack) push(v int) error {

	if s.top == s.MaxTop {
		return errors.New("栈已满")
	}
	s.top++
	s.data[s.top] = v
	return nil
}

func (s *Stack) pop() (v int, e error) {

	if s.top == -1 {
		return -1, errors.New("栈已空")
	}

	v = s.data[s.top]
	s.data[s.top] = 0
	s.top--

	return v, nil
}
