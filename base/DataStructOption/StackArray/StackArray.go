package StackArray

type StackArray interface {
	Clear()             // 清空
	Size() int          // 大小
	Push(v interface{}) // 添加
	Pop() interface{}   // 弹出
	IsFull() bool       // 是否已满
	IsEmpty() bool      // 是否为空
}

//
type Stack struct {
	dataSource  []interface{}
	capSize     int
	currentSize int
}

func NewStack() *Stack {
	s := new(Stack)
	s.dataSource = make([]interface{}, 0, 1000)
	s.capSize = 1000
	s.currentSize = 0
	return s
}

func (s *Stack) Clear() {
	s.currentSize = 0
	s.dataSource = make([]interface{}, 0, 1000)
}

func (s *Stack) Size() int {
	return s.currentSize
}

func (s *Stack) Push(v interface{}) {

	if s.IsFull() {
		return
	}

	s.dataSource = append(s.dataSource, v)
	s.currentSize++
}

func (s *Stack) Pop() interface{} {

	if s.IsEmtpy() {
		return nil
	}

	last := s.dataSource[s.currentSize-1]
	s.dataSource = s.dataSource[:s.currentSize-1] //删除最后一个
	s.currentSize--
	return last
}

func (s *Stack) IsFull() bool {
	return s.currentSize == s.capSize
}

func (s *Stack) IsEmtpy() bool {
	return s.currentSize == 0
}
