package StackArray

/**
f(n)=f(n-1) + f(n-2)
1=1
2=1
3=2
4=3
5=5
6=8
7=13
8=21
9=34
*/
/**
n:代表是第几个数
*/
func Ferbi(n int) int {

	if n == 1 || n == 2 {
		return 1
	}
	return Ferbi(n-1) + Ferbi(n-2)
}

func FerbiStack(n int) int {

	stack := NewStack()
	stack.Push(1)
	stack.Push(1)

	last := 0
	for i := 3; !stack.IsEmtpy(); i++ {

		v1 := stack.Pop().(int)
		v2 := stack.Pop().(int)

		last = v1 + v2
		if i == n {
			break
		}
		//fmt.Println(i, v1, v2, last)
		stack.Push(v1)
		stack.Push(last)
	}

	return last
}

func FerbiStackDesc(n int) int {

	stack := NewStack()
	stack.Push(n)

	last := 0
	for !stack.IsEmtpy() {

		v1 := stack.Pop().(int)

		if v1 == 1 || v1 == 2 {
			last += 1
		} else {
			stack.Push(v1 - 2)
			stack.Push(v1 - 1)
		}
	}

	return last
}
