package src

import (
	"fmt"
	"log"
)

func test2Func() {
	log.Println("Entering function")
	log.Println("ok")
}
func test2hello() {
	// log.Println("hello: Entering function")
	fmt.Println("Hello, world!")
    // more code here
	// log.Println("hello: Exiting function")
}

func test2goodbye() {
	log.Println("goodbye: Entering function")
	fmt.Println("Goodbye, world!")
	// more code here
	// log.Println("goodbye: Exiting function")
}
