package main

import (
	"fmt"
	"os"


)
import 	"github.com/codementor/wman/pkg/string"

func main() {
	name := os.Args[1]
	fmt.Printf("hello %s, your name backward is %q", name, string.Reverse(name))
}
