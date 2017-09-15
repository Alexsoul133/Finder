package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Alexsoul133/111/finder"
)

type Finder interface {
	findcache() []string
	// refreshcache() []string
	newfind() []string
}
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

//поиск в кеш
func (f *Find) findcache() []string {
	var res []string
	cache1, err := finder.Cache(f.cache, f.find)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in Finder.Cache: %s\n", err)
		return nil
	}
	println("Start find in cache")
	for i := range cache1 {
		res = append(res, cache1[i])
	}
	return res
}

//новый поиск и обновление кеша
func (f *Find) newfind() []string {
	var res []string
	find, err := finder.Find(f.path, f.find)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in Finder.Find: %s\n", err)
	}
	println("\nStart new find")
	cache, err := finder.Cache(f.cache, f.find)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in Finder.Cache: %s\n", err)
	}
	cachestr := strings.Join(cache, " ")
	w, err := os.OpenFile(f.cache, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error open cache: %s\n", err)
	}
	for i := range res {
		if !strings.Contains(cachestr, find[i]) {
			_, err = w.WriteString(res[i] + "\n")
		}
	}
	return find
}

func runfinder(f Finder) {
	res := f.findcache()
	for i := range res {
		fmt.Printf("%v\n", res[i])
	}
	res1 := f.newfind()
	for i := range res1 {
		fmt.Printf("%s\n", res1[i])
	}
	return
}

func main() {

	if len(os.Args) < 2 {
		println("Not enough args\nNeed input name and path")
		return
	}
	if len(os.Args) < 3 {
		input := Find{
			cache: "cache.txt",
			path:  "\\" + filepath.Join("\\rackstation3", "YANDEX"),
			find:  os.Args[1]}
		println("Args[3] not input. Will find in \\" + input.path + "\n")
		runfinder(&input)

	} else {
		res, err := finder.Find(os.Args[2], os.Args[1])
		if err != nil {
			println("Error")
		}
		for i := range res {
			fmt.Printf("%s\n", res[i])
		}
	}
	print("\a")
}
