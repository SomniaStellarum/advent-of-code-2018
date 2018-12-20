package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inFile := "input.txt"
	f, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("Cannot open file %v: %v", inFile, err)
	}
	list := make(map[int]bool)
	list_for_sum := make([]int, 0, 1500)
	sum := 0
	list[0] = true
	reader := bufio.NewReader(f)
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		i, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatalf("Could not convert int: %v", err)
		}
		sum += i
		_, ok := list[sum]
		if ok {
			fmt.Println("answer: ", sum)
			return
		}
		list[sum] = true
		list_for_sum = append(list_for_sum, i)
	}
	for {
		for _, l := range list_for_sum {
			sum += l
			_, ok := list[sum]
			if ok {
				fmt.Println("answer: ", sum)
				return
			}
			list[sum] = true
		}
	}
}
