package main

import (
	"fmt"
	"os"
)

func main() {
	name_from_args := os.Args[1]

	fmt.Println("Hello " + name_from_args + "!")
}
