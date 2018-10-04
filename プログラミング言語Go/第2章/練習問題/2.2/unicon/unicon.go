package unicon

func CtoF(c float64) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FtoC(f float64) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func FttoM(ft float64) Metre {
	return Metre(ft / 3.2808)
}
func MtoFt(m float64) Feet {
	return Feet(m * 3.2808)
}
func PtoK(p float64) Kilogram {
	return Kilogram(p / 2.2046)
}
func KtoP(k float64) Pound {
	return Pound(k * 2.2046)
}
