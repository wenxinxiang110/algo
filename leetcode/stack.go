package leetcode

type Stack interface {
	GetTop() int
	Pop() int
	Push(int)
	IsEmpty() bool
}

type stack struct {
	Elem []int
	Top  int
}

func (s *stack) GetTop() int {
	if s.Top >= 0 {
		return s.Elem[s.Top]
	}
	return -1
}

func (s *stack) Pop() (res int) {
	if s.Top >= 0 {
		res = s.Elem[s.Top]
		s.Top--
		s.Elem = s.Elem[:len(s.Elem)-1]
		return
	}
	return -1
}

func (s *stack) Push(num int) {
	s.Top++
	s.Elem = append(s.Elem, num)
}

func (s *stack) IsEmpty() bool {
	return s.Top < 0
}
