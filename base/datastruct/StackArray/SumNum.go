package StackArray

/**
求n以内所有数的和
*/
func Sum(n int) int {

	if n == 0 {
		return n
	}

	return n + Sum(n-1)
}

func SumStack(n int) int {

	stack := NewStack()
	stack.Push(n)
	last := 0
	for !stack.IsEmtpy() {
		v := stack.Pop()

		if v.(int)-1 < 0 {
			break
		} else {
			last += v.(int)
			stack.Push(v.(int) - 1)
		}
	}

	return last
}
