package main

import (
	"fmt"
	"io"
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
	const BufferSize = 100
	file, err := os.Open("filetoread.txt")
	handle := handleFn(file)
	handle(err)

	for {
		// Reinstanting the buffer for each iteration to zero out all the
		// elements. Even this is technically wrong because, say, for a file
		// size of 101 bytes, we'll end up instantiating a 100 element wide
		// slice just to read that one residual byte.
		buffer := make([]byte, BufferSize)
		bytesread, err := file.Read(buffer)

		// We don't have to exit as an error when the value returned as an
		// error is the EOF token.
		if err == io.EOF {
			file.Close()
			break
		} else {
			handle(err)
		}

		fmt.Println("bytes read: ", bytesread)
		fmt.Println("bytestream to string: ", string(buffer))
	}
}
