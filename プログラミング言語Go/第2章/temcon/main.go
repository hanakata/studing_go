package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	const freezingC, boilingF = 0.0, 212.0
	fmt.Println(tempconv.CtoF(freezingC))
	fmt.Println(tempconv.FtoC(boilingF))
}
