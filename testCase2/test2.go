package testCase2

import (
	"fmt"
	"log"
)

func testing() int{
	log.Print("ok")
	fmt.Println("Testing is executed")
	return 2
}

func anotherFunc(x int) {
	log.Print("anotherFunc: Executed")
	fmt.Printf("anotherFunc: %v", x);
}

func main() {
	x := testing()
	anotherFunc(x)
}