package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	fmt.Printf("hello %s!", name)
}
