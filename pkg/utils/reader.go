package utils

import (
	"bufio"
	"os"

	"github.com/vkalekis/advent-of-code/pkg/logger"
)

type Reader interface {
	Read()
	Stream() chan string
}

type FileReader struct {
	StreamCh chan string
	filename string
}

func NewFileReader(filename string) *FileReader {
	return &FileReader{
		StreamCh: make(chan string, 5),
		filename: filename,
	}
}

func (fr *FileReader) Read() {
	defer close(fr.StreamCh)

	file, err := os.Open(fr.filename)
	if err != nil {
		logger.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fr.StreamCh <- scanner.Text()
	}
}

func (fr *FileReader) Stream() chan string {
	return fr.StreamCh
}

type DummyReader struct {
	StreamCh chan string
	filename string
}

func NewDummyReader() *DummyReader {
	return &DummyReader{
		StreamCh: make(chan string, 5),
	}
}

func (dr *DummyReader) Read() {
	defer close(dr.StreamCh)

	dr.StreamCh <- "eighteight9dnvcqznjvfpreight"

}

func (dr *DummyReader) Stream() chan string {
	return dr.StreamCh
}
