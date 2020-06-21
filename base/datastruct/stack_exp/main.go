package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	expFn()
}

func testStack() {
	s := &Stack{top: -1}
	for i := 0; i < len(s.data); i++ {
		s.push(i + 1)
		fmt.Printf("push %d data:%v\n", i, s.data)
	}

	for i := 0; i < len(s.data); i++ {
		v, e := s.pop()
		if e != nil {
			fmt.Println("出栈出错了：" + e.Error())
		} else {
			fmt.Printf("%d, v=%d\n", i, v)
			fmt.Printf("%d, data=%v\n", i, s.data)
		}
	}
}

// 3+2*6-2
// 推算的过程
// 一.
// 数栈：3，2，
// 操作符 + *
// 二.
// 数栈：3，12，2
// 操作符 + -
// 三.
// 数栈：3 10
// 操作符：+
// 四.
// 数栈：13
// 操作符：空
func expFn() {
	// 数栈
	numstack := &Stack{top: -1}

	// 符号栈
	operStack := &Stack{top: -1}

	exp := "345+2*6-20"
	index := 0

	var num, num1, oper, rs int

	keepNum := ""

	for {
		ch := exp[index : index+1] // 单个字符串
		temp := int([]byte(ch)[0]) // 运算符对应的Assii码

		if operStack.isOper(temp) { // 运算符

			if operStack.top == -1 { // 是空栈，直接入栈
				operStack.push(temp)
			} else {
				// 判断优先级
				fmt.Printf("level: %d compare %d\n", operStack.priority(operStack.data[operStack.top]), temp)
				if operStack.priority(operStack.data[operStack.top]) > operStack.priority(temp) { // 获取两个数进行运算

					fmt.Printf("numstack data:%v\n", numstack.data)
					fmt.Printf("operStack data:%v\n", operStack.data)

					num, _ = numstack.pop()
					num1, _ = numstack.pop()
					oper, _ = operStack.pop()
					rs = numstack.cal(num, num1, oper)
					fmt.Printf("%d %s %d=%d\n", num1, numstack.getOper(oper), num, rs)

					// 把计算结果放到数栈中
					numstack.push(rs)
					operStack.push(temp)
				} else {
					operStack.push(temp)
				}

			}
		} else { // 操作数

			// 这一步要转成数字
			//numstack.push(temp)

			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numstack.push(int(val))

			} else { //不在最后一位

				tempch := exp[index+1 : index+2]
				if numstack.isOper(int([]byte(tempch)[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numstack.push(int(val))

					keepNum = ""
				}
			}

			//val, _ := strconv.ParseInt(ch, 10, 64)
			//numstack.push(int(val))
		}

		//循环完成
		if len(exp)-1 == index {
			break
		}

		index++
	}

	for {

		if operStack.top == -1 {
			break
		}

		oper, _ = operStack.pop()
		num, _ = numstack.pop()
		num1, _ = numstack.pop()
		rs = numstack.cal(num, num1, oper)
		fmt.Printf("%d %s %d=%d\n", num1, numstack.getOper(oper), num, rs)

		numstack.push(rs)
	}

	fmt.Println("计算完成！！")
	v, _ := numstack.pop()
	fmt.Printf("结果为：%d\n", v)
}

type Stack struct {
	top  int     // 栈顶的下标,-1代表为空的
	data [20]int // 存放数据的数组
}

func (s *Stack) push(v int) error {

	if len(s.data) == s.top {
		return errors.New("stack is full")
	}

	s.top++
	s.data[s.top] = v

	return nil
}

func (s *Stack) pop() (v int, e error) {

	if s.top == -1 {
		return -1, errors.New("stack is empty")
	}

	v = s.data[s.top]
	s.top--

	return v, nil
}

func (s *Stack) list() {

	for i := s.top; i > -1; i-- {
		fmt.Print(s.data[i])
		fmt.Print(",")
	}
	fmt.Println()
}

// 判断是否一个运算符
func (s *Stack) isOper(v int) bool {

	if v == 42 || v == 43 || v == 45 || v == 47 {
		return true
	}

	return false
}

func (s *Stack) cal(v, v1, oper int) int {
	rs := 0
	switch oper {
	case 42:
		rs = v1 * v
	case 43:
		rs = v1 + v
	case 45:
		rs = v1 - v
	case 47:
		rs = v1 / v
	default:
		fmt.Println("oper is error")
	}

	return rs
}

func (s *Stack) getOper(oper int) string {
	rs := ""
	switch oper {
	case 42:
		rs = "*"
	case 43:
		rs = "+"
	case 45:
		rs = "-"
	case 47:
		rs = "/"
	default:
		fmt.Println("oper is error")
	}

	return rs
}

// 返回操作符的优先级
func (s *Stack) priority(oper int) int {
	rs := 0
	switch oper {
	case 42:
		rs = 1
	case 43:
		rs = 0
	case 45:
		rs = 0
	case 47:
		rs = 1
	default:
		fmt.Println("oper is error")
	}
	return rs
}
