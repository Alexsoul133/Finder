package stack

import "fmt"

type Stack struct {
	stack []string
}

func New() *Stack {
	return &Stack{}
}

//PushBack value to stock
func (s *Stack) PushBack(v string) {
	s.stack = append(s.stack, v)
}

//PopBack value from stock
func (s *Stack) PopBack() (string, error) {
	if len(s.stack) == 0 {
		return "", fmt.Errorf("nothing to pop")
	}
	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res, nil
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) String() string {
	str := "\n"
	for i := range s.stack {
		// str += strconv.Itoa(i) + " " + s.stack[i] + "\n"
		str += fmt.Sprintf("%v %v\n", i, s.stack[i])
	}
	return str
}
