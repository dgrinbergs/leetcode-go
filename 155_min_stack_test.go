package leetcode_go

// https://leetcode.com/problems/min-stack/description/

type MinStack struct {
	values   []int
	minimums []int
}

func Constructor() MinStack {
	return MinStack{
		values:   make([]int, 0),
		minimums: make([]int, 0),
	}
}

func (s *MinStack) Push(value int) {
	s.values = append(s.values, value)

	if len(s.minimums) == 0 {
		s.minimums = append(s.minimums, value)
		return
	}

	minimum := s.minimums[len(s.minimums)-1]
	s.minimums = append(s.minimums, min(value, minimum))
}

func (s *MinStack) Pop() {
	s.values = s.values[:len(s.values)-1]
	s.minimums = s.minimums[:len(s.minimums)-1]
}

func (s *MinStack) Top() int {
	return s.values[len(s.values)-1]
}

func (s *MinStack) GetMin() int {
	return s.minimums[len(s.minimums)-1]
}
