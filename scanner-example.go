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

	// Returns a boolean based on whether there's a next instance of `\n`
	// character in the IO stream. This step also advances the internal pointer
	// to the next position (after '\n') if it did find that token.
	read := scanner.Scan()

	if read {
		fmt.Println("read byte array: ", scanner.Bytes())
		fmt.Println("read string: ", scanner.Text())
	}

	// goto line number 30 and repeat
}
