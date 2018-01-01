package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const BufferSize = 100
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		bytesread, err := file.Read(buffer)

		// err value can be io.EOF, which means that we reached the end of
		// file, and we have to terminate the loop. Note the fmt.Println lines
		// will get executed for the last chunk because the io.EOF gets
		// returned from the Read function only on the *next* iteration, and
		// the bytes returned will be 0 on that read.
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		fmt.Println("bytes read: ", bytesread)
		fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	}
}
