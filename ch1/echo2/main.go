package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 一种更高效的方式: fmt.Println(strings.Join(os.Args[1:], " "))
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
