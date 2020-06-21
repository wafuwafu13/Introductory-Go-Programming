package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius { // kelvin型の k をレシーバとする関数
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	var c celsius

	c = kelvinToCelsius(k)
	c = k.celsius()
}