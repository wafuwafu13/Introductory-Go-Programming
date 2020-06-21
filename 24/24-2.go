package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

func stardate(t time.Time) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

func stardate2(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int (s % 668)
}

// なかったらstardaterインターフェースの要件が満たされない
func (s sol) Hour() int {
	return 0
}

func main() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)

	fmt.Printf("%.1f Curiosity has landed\n", stardate(day)) // 1219.2 Curiosity has landed

	fmt.Printf("%.1f Curiosity has landed\n", stardate2(day)) // 1219.2 Curiosity has landed

	s := sol(1422)

	fmt.Printf("%.1f Happy birthday\n", stardate(s)) // cannot use s (type sol) as type time.Time in argument to stardate
	fmt.Printf("%.1f Happy birthday\n", stardate2(s)) // 1086.0 Happy birthday
}