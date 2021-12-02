package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func transferFileContents(filename string) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		err1 := fmt.Errorf("during file transfer impossible to open source file: %w", err)
		fmt.Println(errors.Is(err1, err))
		return err1
	}
	err = ioutil.WriteFile("/tmp/filecontents", contents, 0644)
	if err != nil {
		return fmt.Errorf("during file transfer impossible to write source file: %w", err)
	}
	return nil
}

func main() {
	err := transferFileContents("/my/imaginary/file")
	if err != nil {
		log.Printf("error occured: %s", err)
	}
}
