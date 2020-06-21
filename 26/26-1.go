package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer) // 0xc000018098

	address := &answer
	fmt.Println(*address) // 42

	fmt.Printf("addressの型は %Tです\n", address) // addressの型は *intです
}