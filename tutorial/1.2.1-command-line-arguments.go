/*
	Run this example:
	go run 1.2-command-line-arguments.go ar1 arg2
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Args:", s)
}
