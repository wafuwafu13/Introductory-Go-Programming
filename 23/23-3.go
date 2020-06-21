package main

import "fmt"

type sol int

type report struct {
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

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	return 5
}

// ambiguous selector report.days を回避
func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	report := report{
		sol: 15,
		temperature: temperature{high: -1.0, low: -78.0},
		location: location{-4.5895, 137.4417},
	}
	fmt.Println(report.days(1446)) // 1431
}