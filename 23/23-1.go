package main

import "fmt"

type report struct {
	sol int
	temperature temperature
	location location
}

type report2 struct {
	sol int
	temperature // 型を埋め込む
	location
}

type sol int

type report3 struct {
	sol
	temperature
	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report1 := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report1) // {sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5895 long:137.4417}}
	fmt.Printf("%v度\n", report1.temperature.high) // -1度
	fmt.Printf("平均%v度\n", t.average()) // 平均-39.5度

	report2 := report2{
		sol: 15,
		temperature: temperature{high: -1.0, low: -78.0},
		location: location{-4.5895, 137.4417},
	}

	fmt.Printf("平均%v度\n", report2.average()) // 平均-39.5度
	fmt.Printf("%v度\n", report2.high) // -1度

	report3 := report3{sol: 15}

	fmt.Println(report3.sol.days(1446)) // 1431
	fmt.Println(report3.days(1446)) // 1431
}