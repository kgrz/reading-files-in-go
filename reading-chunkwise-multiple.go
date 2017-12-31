package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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

	fileinfo, err := file.Stat()
	handle(err)

	filesize := int(fileinfo.Size())
	// Number of go routines we need to spawn.
	concurrency := filesize / BufferSize
	if filesize%BufferSize != 0 {
		concurrency++
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		// buffer size is 100, so we need to read from the 100th byte every
		// iteration
		offset := int64(BufferSize * i)
		go func(f *os.File, offset int64, wg *sync.WaitGroup) {
			buffer := make([]byte, BufferSize)
			// Cheating here again by simply ignoring the error here because we
			// are not in a infinite for loop. This loop will terminate after
			// all the goroutines are done executing the wg.Done function.
			bytesread, _ := file.ReadAt(buffer, offset)

			fmt.Println("bytes read: ", bytesread)
			fmt.Println("bytestream to string: ", string(buffer))
			wg.Done()
		}(file, offset, &wg)
	}

	wg.Wait()
}
