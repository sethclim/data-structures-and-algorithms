package stack

import "fmt"

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) <= 0 {
		return -1, fmt.Errorf("Can not pop from empty stack")
	}

	pop := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return pop, nil
}
