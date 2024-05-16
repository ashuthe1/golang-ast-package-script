package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("main: Starting the application")
	hello()
	goodbye()
	newFunction()
}

func newFunction() {
	log.Println("newFunction: Entering function")
}
func hello() {
	log.Println("hello: Entering function")
	fmt.Println("Hello, world!")
	// more code here
	// log.Println("hello: Exiting function")
}

func goodbye() {
	log.Println("goodbye: Entering function")
	fmt.Println("Goodbye, world!")
	// more code here
	log.Println("goodbye: Exiting function")
}
