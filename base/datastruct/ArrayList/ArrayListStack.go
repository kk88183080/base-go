package ArrayList

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
	dataSource *ArrayList
	capSize    int
}

func NewArrayListStack() *Stack {
	s := new(Stack)
	s.dataSource = NewArrayList()
	s.capSize = 10
	return s
}

func (s *Stack) Clear() {
	s.dataSource.Clear() // 层级调用
	s.capSize = 10
}

func (s *Stack) Size() int {
	return s.dataSource.Size()
}

func (s *Stack) Push(v interface{}) {

	if s.IsFull() {
		return
	}

	s.dataSource.Append(v)
}

func (s *Stack) Pop() interface{} {

	if s.IsEmtpy() {
		return nil
	}

	i := s.dataSource.Size() - 1
	last, _ := s.dataSource.Get(i)
	s.dataSource.Delete(i)

	return last
}

func (s *Stack) IsFull() bool {
	return s.dataSource.Size() >= s.capSize
}

func (s *Stack) IsEmtpy() bool {
	return s.dataSource.Size() == 0
}
