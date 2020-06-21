package main

import (
	"fmt"
	"strings"
)

type martian struct{}

type creater struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

type talker interface {
	talk() string
}

type starship struct {
	laser // 型を埋め込む
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	var t interface {
		talk() string
	}

	t = martian{}
	fmt.Println(t.talk()) // nack nack

	t = laser(3) // 3をlaser型に変換
	fmt.Println(t.talk()) // pewpewpew

	shout(martian{}) // NACK NACK
	shout(laser(2)) // PEWPEW

	shout(creater{}) // creater does not implement talker (missing talk method)

	s := starship{laser(3)}

	fmt.Println(s.talk()) // pewpewpew
	shout(s) // PEWPEWPEW
}