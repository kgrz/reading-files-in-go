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
	scanner.Split(bufio.ScanWords)

	// initial size of our wordlist
	bufferSize := 10
	words := make([]string, bufferSize)
	pos := 0

	for scanner.Scan() {
		words[pos] = scanner.Text()
		pos++

		if pos >= len(words) {
			// expand the buffer by 100 again
			newbuf := make([]string, bufferSize)
			words = append(words, newbuf...)
		}
	}

	fmt.Println("word list:")
	for _, word := range words {
		fmt.Println(word)
	}
}
