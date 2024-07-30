package datastructures

type Stack[T any] struct {
	stack []*T
}

func (s *Stack[T]) Push(value *T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Pop() (*T, bool) {
	if len(s.stack) == 0 {
		return nil, false
	}

	last := len(s.stack) - 1
	value := s.stack[last]
	s.stack = s.stack[:last]

	return value, true
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}
