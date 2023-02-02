package main

import (
	"fmt"
	"os"
)

func main() {
	// Get args without the path
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please provide a name when callin me!")
		fmt.Println("example:	./2-CMA2.exe Alex")
	} else {
		for i := 0; i < len(args); i++ {
			fmt.Println("Hello " + args[i] + "!")
		}
	}
}
