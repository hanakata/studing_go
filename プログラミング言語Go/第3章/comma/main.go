package main

import (
	"flag"
	"fmt"

	"./comma"
)

var s = flag.String("s", "100000", "string flag")

func main() {
	flag.Parse()
	dir := *s
	fmt.Println(comma.Comma(dir))
}
