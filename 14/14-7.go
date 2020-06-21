package main

import "fmt"

type kelvin float64

func main() {
	var k kelvin = 294.0
	sensor := func() kelvin { // 無名関数はスコープにある変数を囲い込む
		return k
	}

	fmt.Println(sensor()) // 294

	k++

	fmt.Println(sensor()) // 295
}