package stack

type Value struct {
	res []string
}

type Stack struct {
	stack []*Value
	path  string
	cap   int
}

// func New(s *Stack) []string {
// 	s.stack = make([]string, len(s.stack))
// 	return s.stack
// }

//Push value to stock
func (s *Stack) Push(v *Value) {
	stack := make([]*Value, len(v.res))
	copy(s.stack, stack)
	s.stack = stack
	s.cap++
	return
}

//Pop value from stock
func (s *Stack) Pop() *Value {
	if s.cap <= 0 {
		return nil
	}
	stack := s.stack[s.cap-1]
	s.cap--
	return stack
}
