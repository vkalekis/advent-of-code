package utils

import (
	"bufio"
	"log"
	"os"
)

type Reader interface {
	Read()
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
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fr.StreamCh <- scanner.Text()
	}
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
