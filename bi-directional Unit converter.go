package main

import (
	"fmt"; "math"
)

const (
	pi = math.Pi
)

//Converter to change unit
type Converter struct{}

type (
	Minutes      float64
	Seconds      float64
	Milliseconds float64
	Feet         float64
	Centimeter   float64
	Celsius      float64
	Fahrenheit   float64
	Radian       float64
	Degree       float64
	Kilogram     float64
	Pounds       float64
)

//CentimeterToFeet attaching to Converter struct

func (cvr Converter) SecondsToMilliseconds(s Seconds) Milliseconds {
	Ms := Milliseconds(float64(s) * 1000)
	return Ms
}
func (cvr Converter) MillisecondsToSeconds(ms Milliseconds) Seconds {
	sec := Seconds(float64(ms) / 1000)
	return sec
}
func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	cm := Centimeter(float64(f) * 30.48)
	return cm
}
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	ft := Feet(float64(c) / 30.48)
	return ft
}
func (cvr Converter) CelsiusToFahrenheit(C Celsius) Fahrenheit {
	Fahr := Fahrenheit(((float64(C) * 9 )/ 5) + 32)
	return Fahr
}
func (cvr Converter) FahrenheitToCelsius(F Fahrenheit) Celsius {
	Cel := Celsius(((float64(F) - 32) * 5) / 9)
	return Cel
}
func (cvr Converter) RadianToDegree(rad Radian) Degree {
	Deg := Degree((float64(rad) * 180) / pi)
	return Deg
}
func (cvr Converter) DegreeToRadian(deg Degree) Radian {
	Rad := Radian((float64(deg) * pi) / 180)
	return Rad
}
func (cvr Converter) KilogramToPounds(kg Kilogram) Pounds {
	Pnd := Pounds(float64(kg) * 2.205)
	return Pnd
}
func (cvr Converter) PoundsToKilogram(Pnd Pounds) Kilogram {
	kg := Kilogram(float64(Pnd) /2.205)
	return kg
}

func main() {

	cvr := Converter{}
	fmt.Println(cvr.CentimeterToFeet(180))
	fmt.Println(cvr.FeetToCentimeter(6))
	fmt.Println(cvr.SecondsToMilliseconds(60))
	fmt.Println(cvr.MillisecondsToSeconds(6))
	fmt.Println(cvr.CelsiusToFahrenheit(100))
	fmt.Println(cvr.FahrenheitToCelsius(212))
	fmt.Println(cvr.RadianToDegree(pi))
	fmt.Println(cvr.DegreeToRadian(180))
	fmt.Println(cvr.KilogramToPounds(1))
	fmt.Println(cvr.PoundsToKilogram(2.205))

}
