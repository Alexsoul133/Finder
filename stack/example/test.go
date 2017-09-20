package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	finder "github.com/Alexsoul133/Finder"
)

type Value struct {
	res string
}

type Stack struct {
	stack []*Value
	cap   int
}
type Find struct {
	path  string
	find  string
	cache string
}

const cachefile = "cac.che"

func newfind(path string, find string) *Find {
	return &Find{
		cache: cachefile,
		path:  path,
		find:  find,
	}
}

//Push value to stock
func (s *Stack) Push(v *Value) {
	stack := make([]*Value, len(v.res))
	copy(s.stack, stack)
	s.stack = stack
	s.stack[s.cap] = v
	s.cap++
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

func (f *Find) newfind() []string {
	find, err := finder.Find(f.path, ".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in Finder.Find: %s\n", err)
	}
	println("\nStart new find")
	return find
}

func runfinder(f *Find) {
	res := f.newfind()

	s := &Stack{stack: make([]*Value, len(res))}

	for i := range res {
		s.Push(&Value{res[i]})
	}
	re, _ := regexp.Compile(f.find)
	e := []string{}
	for i := range res {
		e[i] = s.Pop().res
		if re.MatchString(e[i]) {
			fmt.Printf("%s\n", e[i])
		}
	}
	return
}

func main() {
	input := newfind("\\"+filepath.Join("\\rackstation3", "YANDEX"), os.Args[1])
	println("Args[3] not input. Will find in " + input.path + "\n")
	runfinder(input)

	// s := &Stack{stack: make([]*Value, 3)}
	// s.Push(&Value{"1"})
	// fmt.Printf("%v", s.Pop().res)
}
