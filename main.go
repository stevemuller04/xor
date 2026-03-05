package main

import (
	"bufio"
	"bytes"
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
	readers := make([]CyclicInput, n)
	for i := range n {
		if f, err := os.Open(os.Args[i+1]); err != nil {
			log.Fatal(err)
			return
		} else {
			defer f.Close()
			readers[i] = NewInput(f)
		}
	}

	// XOR all the bytes
	for {
		var result byte = 0
		var allEOF bool = true

		for i := range n {
			if err := readers[i].ReadNextByte(); err != nil {
				log.Fatal(err)
			} else if b, ok := readers[i].GetByte(); ok {
				result ^= b
			}
			allEOF = allEOF && readers[i].EOF
		}

		if allEOF {
			break
		} else {
			f.WriteByte(result)
		}
	}
}

type CyclicInput struct {
	reader io.ReadSeeker
	buffer bytes.Buffer
	EOF    bool
}

func NewInput(f io.ReadSeeker) CyclicInput {
	return CyclicInput{f, bytes.Buffer{}, false}
}

func (i *CyclicInput) cyclicRead(data []byte) (int, error) {
	if n, err := i.reader.Read(data); n > 0 {
		return n, err
	} else if err != nil && err != io.EOF {
		return 0, err
	} else {
		i.EOF = true
		if _, err := i.reader.Seek(0, io.SeekStart); err != nil {
			return 0, err
		} else {
			return i.reader.Read(data)
		}
	}
}

func (i *CyclicInput) ReadNextByte() error {
	if i.buffer.Len() != 0 {
		return nil // we still have data
	}

	var data [4096]byte
	if n, err := i.cyclicRead(data[:]); err != nil && err != io.EOF {
		return err
	} else {
		_, err := i.buffer.Write(data[:n])
		return err
	}
}

func (i *CyclicInput) GetByte() (b byte, ok bool) {
	if b, err := i.buffer.ReadByte(); err != nil {
		return 0, false
	} else {
		return b, true
	}
}
