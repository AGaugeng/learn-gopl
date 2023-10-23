package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countFiles := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v/n", err)
				continue
			}
			countLines(f, counts, countFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, countFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, countFlies map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		countFlies[text] = append(countFlies[text], f.Name())
	}
}
