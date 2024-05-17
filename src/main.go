package src

import (
	"fmt"
	"log"
)

func testFunc() {
	log.Println("testFunc: Entering function")
	log.Println("testFunc: ok")
	fmt.Print("testFunc: Hehe")
}
func testhello() {
	log.Println("testhello: Entering function")
	fmt.Println("testhello, world!")
    // more code here
	// log.Println("hello: Exiting function")
}

func testgoodbye() {
	log.Println("goodbye: Entering function")
	fmt.Println("Goodbye, world!")
	// more code here
	// log.Println("goodbye: Exiting function")
}

func main() {
	log.Println("Starting the application")
	testFunc()
	testhello()
	testgoodbye()
}