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
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Args:", s)
}
