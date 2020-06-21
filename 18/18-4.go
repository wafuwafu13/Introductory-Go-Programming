package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: 長さ %v, 容量 %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	planets2 := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestiral := planets[0:4:4]
	dump("terrestiral", terrestiral) // terrestiral: 長さ 4, 容量 4 [Mercury Venus Earth Mars]

	worlds := append(terrestiral, "ケレス") 
	dump("worlds", worlds) // worlds: 長さ 5, 容量 8 [Mercury Venus Earth Mars ケレス]

	dump("planets", planets) // planets: 長さ 8, 容量 8 [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]

	terrestiral2 := planets2[0:4]
	dump("terrestiral2", terrestiral2) // terrestiral2: 長さ 4, 容量 8 [Mercury Venus Earth Mars]

	worlds2 := append(terrestiral2, "ケレス")
	dump("worlds2", worlds2) // worlds2: 長さ 5, 容量 8 [Mercury Venus Earth Mars ケレス]

	dump("planets2", planets2) // planets2: 長さ 8, 容量 8 [Mercury Venus Earth Mars ケレス Saturn Uranus Neptune]


}