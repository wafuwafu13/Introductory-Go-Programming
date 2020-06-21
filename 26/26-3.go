package main

import "fmt"

type person struct {
	name, superpower string
	age int
}

type person2 struct {
	name string
	age int
}

func birthday(p *person) {
	p.age++ // (*p).age++
}

// レシーバがポイント型じゃなかったらTerryは15歳のまま
func (p *person2) birthday2() {
	p.age++
}

type stats struct {
	level int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name string
	stats stats
}

func main() {
	rebecca := person{
		name: "Rebecca",
		superpower: "imagination",
		age: 14,
	}

	birthday(&rebecca)

	fmt.Printf("%+v\n", rebecca) // {name:Rebecca superpower:imagination age:15}

	terry := &person2{
		name: "Terry",
		age: 15,
	}

	terry.birthday2()

	fmt.Printf("%+v\n", terry) // &{name:Terry age:16}

	nathan := person2{
		name: "Nathan",
		age: 17,
	}

	nathan.birthday2()

	fmt.Printf("%+v\n", nathan) // {name:Nathan age:18}

	player := character{name: "Matthias"}
	levelUp(&player.stats) // 構造体の中にあるフィールドを指し示す
	fmt.Printf("%+v\n", player) // {name:Matthias stats:{level:1 endurance:56 health:280}}
}