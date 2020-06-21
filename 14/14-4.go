package main

import "fmt"

func main() {
	f := func(message string) { // 無名関数を変数に代入
		fmt.Println(message) // Go to the party
	}
	f("Go to the party")
}