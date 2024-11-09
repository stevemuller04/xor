package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	n := len(os.Args) - 1
	if n < 1 {
		fmt.Printf("Usage: %s <file1> <file2> ...\n", os.Args[0])
		return
	}

	// Open files for reading
	readers := make([]*bufio.Reader, n)
	for i := 0; i < n; i++ {
		if f, err := os.Open(os.Args[i + 1]); err != nil {
			log.Fatal(err)
			return
		} else {
			defer f.Close()
			readers[i] = bufio.NewReader(f)
		}
	}

	// XOR all the bytes
	for {
		var result byte = 0
		for i := 0; i < n; i++ {
			if b, err := readers[i].ReadByte(); err == io.EOF {
				return
			} else if err != nil {
				log.Fatal(err)
				return
			} else {
				result ^= b
			}
		}
		f.WriteByte(result)
	}
}
