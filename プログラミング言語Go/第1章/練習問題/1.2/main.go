package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Print(i)
		fmt.Println(" " + arg)
	}
}
