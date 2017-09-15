package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	var res []string
	res = append(res, "111")
	fmt.Printf(res[0])
	fmt.Fprintf(file, res[0])
}
