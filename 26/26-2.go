package main

import "fmt"

type person struct {
	name, superpower string
	age int
}

func main() {
	var administrator *string
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator) // Christopher J. Scolese

	bolden := "Charles F. Bolden"
	administrator = &bolden // 以降、administratorはboldenを指し続ける
	fmt.Println(administrator) // 0xc000010220
	fmt.Println(*administrator) // Charles F. Bolden

	bolden = "Charles F. Bolden Jr."
	fmt.Println(*administrator) // Charles F. Bolden Jr.

	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(*administrator) // Maj. Gen. Charles Frank Bolden Jr.

	major := administrator
	*major = "Major Genera Charles Flank Bolden Jr."
	fmt.Println(bolden) // Major Genera Charles Flank Bolden Jr.

	fmt.Println(administrator) // 0xc000010220
	fmt.Println(administrator == major) // ture

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot // administratorはlightfootを指すことになった
	fmt.Println(administrator == major) // false

	charles := *major // 文字列のコピーが生じる
	*major = "Charles Bolden"

	fmt.Println(charles) // Major Genera Charles Flank Bolden Jr.
	fmt.Println(bolden) // Charles Bolden

	charles = "Charles Bolden"

	fmt.Println(charles == bolden) // true
	fmt.Println(&charles == &bolden) // false

	timmy := &person{
		name: "Timothy",
		age: 10,
	}

	timmy.superpower = "flying" // (*timmy).superpowerと同義
	fmt.Printf("%+v\n", timmy) // &{name:Timothy superpower:flying age:10}
}