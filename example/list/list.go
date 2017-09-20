package main

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	finder "github.com/Alexsoul133/Finder"
)

const cachefile = "cac.che"

type Find struct {
	path  string
	find  string
	cache string
}

type Video struct {
	name    string
	coid    string
	year    string
	quality int
	widht   int
	height  int
	format  string
}

func newfind(path string, find string) *Find {
	return &Find{
		cache: cachefile,
		path:  path,
		find:  find,
	}
}

//новый поиск
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
	stack := list.New()
	for i := range res {
		stack.PushBack(res[i])
	}
	re, _ := regexp.Compile(f.find)
	for ins := stack.Front(); ins != nil; ins = ins.Next() {
		if re.MatchString(ins.Value.(string)) {
			fmt.Printf("%s\n", ins.Value)
		}
	}
	return
}

func main() {

	if len(os.Args) < 2 {
		println("Not enough args\nNeed input find name")
		return
	}
	if len(os.Args) < 3 {
		input := newfind("\\"+filepath.Join("\\rackstation3", "YANDEX"), os.Args[1])
		println("Args[3] not input. Will find in " + input.path + "\n")
		runfinder(input)

	} else {
		input := newfind(os.Args[2], os.Args[1])
		println("Will find in " + input.path + "\n")
		runfinder(input)
	}
	print("\a")
}
