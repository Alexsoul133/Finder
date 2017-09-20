package main

import "fmt"

type stack []string

type res struct {
	stack []*stack
	len   int
}

func (r *res) Push(v *stack) {
	s := append(r.stack, v)
	return
}

func (r *res) Pop() []string {

}

func main() {
	e := make(stack, 0)
	e = e.Push("1")
	e = e.Push("2")
	fmt.Printf("%v,%v", e.Pop(), e.Pop())
}
