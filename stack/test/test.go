package main

import (
	"fmt"

	"github.com/Alexsoul133/Finder/stack"
)

func main() {
	s := stack.New()
	s.PushBack("first")
	s.PushBack("second")
	s.PushBack("third")
	fmt.Print(s)
	fmt.Printf("%v\n", s.Len())
	s.PopBack()
	s.PopBack()
	s.PopBack()
	fmt.Printf("%v\n", s.Len())
	s.PopBack()

}
