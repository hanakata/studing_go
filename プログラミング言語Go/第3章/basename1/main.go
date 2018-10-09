package main

import (
	"flag"
	"fmt"

	"./basename"
)

var s = flag.String("s", "/example/sample.go", "string flag")

func main() {
	flag.Parse()
	dir := *s
	fmt.Println(basename.Basename(dir))
}
