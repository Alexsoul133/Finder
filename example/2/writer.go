package main

import (
	"fmt"
	"os"
)

func main() {
	// file := "cache1.txt"
	// if _, err := os.Stat(file); os.IsExist(err) {
	// 	println("Cache not exist")
	// 	_, err := os.Create(file)
	// 	if err != nil {
	// 		println("Error create file")
	// 	}
	// }
	var res []string
	f, err := os.Create("cache1.txt")
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	defer f.Close()
	// w := bufio.NewWriter(f)
	res = append(res, "1221")
	println(res[0])
	fmt.Fprintf(f, res[0])
	// n3, err := f.WriteString(res[0])
	// fmt.Printf("wrote %d bytes\n", n3)

}
