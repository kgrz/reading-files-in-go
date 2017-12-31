package main

import (
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

	fileinfo, err := file.Stat()
	handle(err)

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	handle(err)

	fmt.Println("bytes read: ", bytesread)
	fmt.Println("bytestream to string: ", string(buffer))
}
