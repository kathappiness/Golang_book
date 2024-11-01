package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dups := make(map[string]map[string]int)
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			line := input.Text()
			counts[line]++
			if dups[line] == nil {
				dups[line] = make(map[string]int)
			}
			dups[line][arg]++
		}
		f.Close()
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for filename, counter := range dups[line] {
				fmt.Printf("\t%d,%v\n", counter, filename)
			}
		}
	}
}
