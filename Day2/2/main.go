package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inFile := "input.txt"
	f, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("Cannot open file %v: %v", inFile, err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var list []string
	for s.Scan() {
		list = append(list, s.Text())
	}

	for i, l1 := range list {
		for _, l2 := range list[i+1:] {
			l_diff, oneDiff := OneDiff(l1, l2)
			if oneDiff {
				fmt.Println("answer: ", l_diff)
				return
			}
		}
	}
}

func OneDiff(l1, l2 string) (string, bool) {
	idx := -1
	if len(l1) != len(l2) {
		return "", false
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] == l2[i] {
			continue
		}
		if idx >= 0 {
			return "", false
		}
		idx = i
	}

	return l1[:idx] + l1[idx+1:], true
}
