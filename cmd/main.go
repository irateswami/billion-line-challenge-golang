package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

const (
	exitFail      = 1
	fileChunkSize = 1024 * 1024
)

type processor struct {
	wg  *sync.WaitGroup
	sm  *sync.Map
	mut *sync.RWMutex
	c   *atomic.Uint64
}

func run() error {
	p := processor{
		wg:  &sync.WaitGroup{},
		mut: &sync.RWMutex{},
		c:   &atomic.Uint64{},
	}

	count := 0
	path := os.Getenv("BILLION_LINE_FILE")
	if path == "" {
		return errors.New("no path found")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	buf := make([]byte, fileChunkSize)
	readStart := 0

	reader := bufio.NewReader(file)
	for {
		n, err := reader.Read(buf[readStart:])
		if err != nil {
			if err == io.EOF {
				break
			}

			return fmt.Errorf("error has occurred: %s", err)
		}

		if readStart+n == 0 {
			break
		}

		chunk := buf[:readStart+n]

		lastNewLine := bytes.LastIndexByte(chunk, '\n') + 1
		bufCopy := make([]byte, lastNewLine)
		copy(bufCopy, buf)

		p.wg.Add(1)
		go p.processBuf(bufCopy)

		count++
	}
	p.wg.Wait()

	fmt.Printf("count: %d\n", p.c.Load())

	return nil
}

func (p *processor) processBuf(buf []byte) {
	var count uint64 = 0

	for _, b := range buf {
		if b == '\n' {
			count++
		}
	}

	p.c.Add(count)

	p.wg.Done()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}
