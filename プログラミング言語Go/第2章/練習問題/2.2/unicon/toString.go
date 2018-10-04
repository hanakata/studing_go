package unicon

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Metre float64
type Kilogram float64
type Pound float64

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%gFt", ft)
}
func (m Metre) String() string {
	return fmt.Sprintf("%gM", m)
}
func (kg Kilogram) String() string {
	return fmt.Sprintf("%gkg", kg)
}
func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}
