package finder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type video struct {
	name    string
	coid    string
	year    string
	quality int
	widht   int
	height  int
	format  string
}

func Find(path string, find string) ([]string, error) {
	var res []string
	find = strings.ToLower(find)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Readdir: %s\n", err)
		return nil, err
	}
	for i := range files {
		if strings.Contains(strings.ToLower(files[i].Name()), find) && !files[i].IsDir() {
			res = append(res, strings.ToLower(path+"\\"+files[i].Name()))
		} else {
			if files[i].IsDir() {
				res1, err := Find(path+"\\"+files[i].Name(), find)
				res = append(res, res1...)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error in FindDir: %s\n", err)
				}
			}
		}
	}
	return res, err
}

// func Find(path string, find string) ([]string, error) {
// 	var res []string
// 	find = strings.ToLower(find)
// 	f, _ := os.Open(path)
// 	files, err := f.Readdirnames(0)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error Readdir: %s\n", err)
// 		return nil, err
// 	}
// 	for i := range files {
// 		if strings.Contains(strings.ToLower(files[i]), find) && strings.Contains(strings.ToLower(files[i]), ".") {
// 			res = append(res, strings.ToLower(path+"\\"+files[i]))
// 		} else {
// 			if !strings.Contains(strings.ToLower(files[i]), ".") {
// 				res1, err := Find(path+"\\"+files[i], find)
// 				res = append(res, res1...)
// 				if err != nil {
// 					fmt.Fprintf(os.Stderr, "Error in FindDir: %s\n", err)
// 				}
// 			}
// 		}
// 	}
// 	return res, err
// }

func Cache(name string, find string) ([]string, error) {
	res := []string{}
	find = strings.ToLower(find)

	if _, err := os.Stat(name); err != nil {
		fmt.Printf("%v not found. Create \n", name)
		os.Create(name)
	}
	fmt.Printf("Found %v \n", name)
	w, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	s1 := strings.Split(string(w), "\n")
	for i := range s1 {
		if strings.Contains(filepath.Base(s1[i]), find) {
			res = append(res, s1[i])
		}
	}

	return res, nil
}
