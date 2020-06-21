package main

import "fmt"

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "未設定"
	}
	return fmt.Sprintf("%d です", n.value)
}

func main() {
	n := newNumber(42)
	fmt.Println(n) // 42です

	e := number{}
	fmt.Println(e) // 未設定

	p := 3
	fmt.Println(p) // 3
}