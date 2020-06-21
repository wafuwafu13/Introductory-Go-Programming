package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: 長さ %v, 容量 %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs1", dwarfs1) // dwarfs1: 長さ 5, 容量 5 [Ceres Pluto Haumea Makemake Eris]

	dwarfs2 := append(dwarfs1, "Orcus") // dwarfs1の内容を、2倍の容量を持つ新規に割り当てた配列にコピーする
	dump("dwarfs2", dwarfs2) // dwarfs2: 長さ 6, 容量 10 [Ceres Pluto Haumea Makemake Eris Orcus]

	dwarfs3 := append(dwarfs2, "Sakacia", "Quaoar", "Sedna")
	dump("dwarfs3", dwarfs3) // dwarfs3: 長さ 9, 容量 10 [Ceres Pluto Haumea Makemake Eris Orcus Sakacia Quaoar Sedna]

	dwarfs3[1] = "Pluto!"

	dump("dwarfs1", dwarfs1) // [Ceres Pluto Haumea Makemake Eris]
	dump("dwarfs2", dwarfs2) // [Ceres Pluto! Haumea Makemake Eris Orcus]
	dump("dwarfs3", dwarfs3) // [Ceres Pluto! Haumea Makemake Eris Orcus Sakacia Quaoar Sedna]
}