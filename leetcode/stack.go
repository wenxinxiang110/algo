package leetcode

import "errors"

var (
	StackEmptyError = errors.New("MyStack is empty")
)

type Stack interface {
	GetTop() (interface{}, error)
	Pop() (interface{}, error)
	Push(interface{})
	IsEmpty() bool
}

type MyStack struct {
	Elem []interface{}
}

func (s *MyStack) GetTop() (interface{}, error) {
	if l := len(s.Elem); l > 0 {
		return s.Elem[l-1], nil
	}
	return nil, StackEmptyError
}

func (s *MyStack) Pop() (res interface{}, err error) {
	if l := len(s.Elem); l > 0 {
		res = s.Elem[l-1]
		s.Elem = s.Elem[:l-1]
		return
	}
	return nil, StackEmptyError
}

func (s *MyStack) Push(num interface{}) {
	s.Elem = append(s.Elem, num)
}

func (s *MyStack) IsEmpty() bool {
	return len(s.Elem) == 0
}
