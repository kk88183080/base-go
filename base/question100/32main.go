package main

import "fmt"

func main() {
	//_32test()
	//_32test1()
	//_32test2()
	_32countTest()
}

type _32foo struct {
	bar string
}

// 错误的写法
func _32test() {
	s1 := []_32foo{
		{"A"},
		{"B"},
		{"C"},
	}

	s2 := make([]*_32foo, len(s1))

	for i, value := range s1 {

		s2[i] = &value

	}

	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}

// 正确的写法1
func _32test1() {
	s1 := []_32foo{
		{"A"},
		{"B"},
		{"C"},
	}

	s2 := make([]*_32foo, len(s1))

	for i, value := range s1 {
		v := value
		s2[i] = &v
	}

	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}

// 正确的写法2
func _32test2() {
	s1 := []_32foo{
		{"A"},
		{"B"},
		{"C"},
	}

	s2 := make([]*_32foo, len(s1))

	for i, _ := range s1 {
		s2[i] = &s1[i]
	}

	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}

func _32countTest() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	count := 0
	for k, v := range m {
		if count == 0 {
			delete(m, "A")
		}

		count++
		fmt.Println(k, v)
	}

	fmt.Printf("count: %d\n", count)
}
