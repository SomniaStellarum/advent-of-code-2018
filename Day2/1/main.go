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
	reader := bufio.NewReader(f)
	sums := make([]int, 2)
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		ts := TwosAndThrees(b)
		sums[0] += ts[0]
		sums[1] += ts[1]
	}
	fmt.Println("answer: ", sums[0]*sums[1])
}

func TwosAndThrees(b []byte) []int {
	m := make(map[byte]int)
	for _, by := range b {
		_, ok := m[by]
		if ok {
			m[by] += 1
		} else {
			m[by] = 1
		}
	}
	ts := make([]int, 2)
	for _, v := range m {
		if v == 2 {
			ts[0] = 1
		}
		if v == 3 {
			ts[1] = 1
		}
	}
	return ts
}
