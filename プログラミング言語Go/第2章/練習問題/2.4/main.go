package main

import (
	"flag"
	"fmt"

	"./popcount"
)

var u = flag.Uint64("u", 32, "uint flag")

func main() {
	flag.Parse()
	check := *u
	fmt.Println(popcount.PopCount(check))
}
