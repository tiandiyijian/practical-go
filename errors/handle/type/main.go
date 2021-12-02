package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type ReadingError struct {
	IOError  error
	Filename string
}

type WritingError struct {
	IOError  error
	Filename string
}

func (e *ReadingError) Error() string {
	return fmt.Sprintf("an error occured while attempting to read the file %s", e.Filename)
}

func (e *ReadingError) Unwrap() error {
	return e.IOError
}

func (e *WritingError) Error() string {
	return fmt.Sprintf("an error occured while attempting to read the file %s", e.Filename)
}

func (e *WritingError) Unwrap() error {
	return e.IOError
}

func transferFileContents(filename string) error {
	contents, err := ioutil.ReadFile(filename)
	// fmt.Println(err)
	if err != nil {
		return &ReadingError{err, filename}
	}
	err = ioutil.WriteFile("/tmp/filecontents", contents, 0644)
	if err != nil {
		return fmt.Errorf("during file transfer impossible to write source file: %w", err)
	}
	return nil
}

func main() {
	err := transferFileContents("/my/imaginary/file")
	var readingError *ReadingError
	if errors.As(err, &readingError) {
		// 我觉得使用指针的指针的原因是要改变原指针的值
		log.Fatalf("error of reading occured: %s: %s", readingError, readingError.Unwrap())
	}
	var writingError *WritingError
	if errors.As(err, &writingError) {
		log.Fatalf("error of writing occured: %s", err)
	}
	log.Println("transfer done")

}
