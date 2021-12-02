package main

import (
	config "error-example/config3"
	"fmt"
	"log"
)

func main() {
	confData, err := config.Load()
	// fmt.Printf("%T\n", err)
	var errinterface error = err
	fmt.Printf("%T\n\n", errinterface)
	if ev, ok := err.(error); ok {
		fmt.Printf("err type: %T ev type: %T\n", err, ev)
		fmt.Println(err == ev)
	} else {
		fmt.Println(ev, ok)
	}
	if err != nil {
		log.Fatalf("Impossible to load application config because: %s", err)
	}
	fmt.Println(confData)
}
