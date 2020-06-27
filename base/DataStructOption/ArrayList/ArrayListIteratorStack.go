package ArrayList

type StackArrayX interface {
	Clear()             // 清空
	Size() int          // 大小
	Push(v interface{}) // 添加
	Pop() interface{}   // 弹出
	IsFull() bool       // 是否已满
	IsEmpty() bool      // 是否为空
}

//
type StackX struct {
	dataSource *ArrayList
	Myit       Iterator
}

func NewArrayListStackX() *StackX {
	s := new(StackX)
	s.dataSource = NewArrayList()
	s.Myit = s.dataSource.Iterator()
	return s
}

func (s *StackX) Clear() {
	s.dataSource.Clear() // 层级调用
	s.dataSource.theSize = 0
}

func (s *StackX) Size() int {
	return s.dataSource.Size()
}

func (s *StackX) Push(v interface{}) {

	if s.IsFull() {
		return
	}

	s.dataSource.Append(v)
}

func (s *StackX) Pop() interface{} {

	if s.IsEmtpy() {
		return nil
	}

	i := s.dataSource.Size() - 1
	last, _ := s.dataSource.Get(i)
	s.dataSource.Delete(i)

	return last
}

func (s *StackX) IsFull() bool {
	return s.dataSource.Size() >= 10
}

func (s *StackX) IsEmtpy() bool {
	return s.dataSource.Size() == 0
}
