package main

import "fmt"

func main() {
	func() {
		fmt.Println("Functions anonymous")
	}() // 無名関数を即時実行
}