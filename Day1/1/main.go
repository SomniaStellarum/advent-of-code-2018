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
	sum := 0
	reader := bufio.NewReader(f)
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(sum)
			log.Fatalf("Could not read line: %v", err)
		}
		i, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatalf("Could not convert int: %v", err)
		}
		sum += i
	}
}
