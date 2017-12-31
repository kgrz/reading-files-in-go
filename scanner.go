package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func handleFn(file *os.File) func(error) {
	return func(err error) {
		if err != nil {
			file.Close()
			log.Fatal(err)
		}
	}
}

func main() {
	file, err := os.Open("filetoread.txt")
	handle := handleFn(file)
	handle(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// This is our buffer now
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("read lines:")
	for _, line := range lines {
		fmt.Println(line)
	}
}
