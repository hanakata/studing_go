package main

import (
	"flag"
	"fmt"

	"./unicon"
)

var (
	t = flag.Float64("t", 32.5, "temp flag")
	l = flag.Float64("l", 100.5, "length flag")
	w = flag.Float64("w", 60.5, "weight flag")
)

func main() {
	flag.Parse()
	temp := *t
	length := *l
	weight := *w
	fmt.Printf("%s = %s\n", unicon.Fahrenheit(temp), unicon.FtoC(temp))
	fmt.Printf("%s = %s\n", unicon.Celsius(temp), unicon.CtoF(temp))
	fmt.Printf("%s = %s\n", unicon.Feet(length), unicon.FttoM(length))
	fmt.Printf("%s = %s\n", unicon.Metre(length), unicon.MtoFt(length))
	fmt.Printf("%s = %s\n", unicon.Kilogram(weight), unicon.PtoK(weight))
	fmt.Printf("%s = %s\n", unicon.Pound(weight), unicon.KtoP(weight))
}
