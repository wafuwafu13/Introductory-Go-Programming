package main

import "fmt"

type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor { // sensor関数型の無名関数を返す
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor()) // 5
}