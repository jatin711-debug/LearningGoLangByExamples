package main

import (
	"fmt"
)

const (
	sunday = iota
	monday
)

type Animals interface {
	walk()
	bark()
}

type Human struct {
	legs  int
	hands int
}

type Deer struct {
	legs int
}

func (h Human) walk() {
	fmt.Println("Human HJere", h.legs)
}

func (h Human) bark() {
	fmt.Println("Human Speech")
}

func (d Deer) walk() {
	fmt.Println("Deer walk", d.legs)
}

func (d Deer) bark() {
	fmt.Println("Deer bark")
}

func perform(a Animals) {
	a.walk()
	a.bark()
}

func main() {
	h := Human{2, 2}
	d := Deer{4}
	perform(h)
	perform(d)
}
