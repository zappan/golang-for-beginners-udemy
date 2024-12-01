package types

import (
	"fmt"
	"myapp/util"
)

// interface type
type Pet interface {
	says() string
	legs() int
}


type Dog struct {
	Name string
	Sound string
	NumLegs int
}

type Cat struct {
	Name string
	Sound string
	NumLegs int
	HasTail bool
}

// NO WAY TO SAY Cat & Dog ARE Pet TO IMPLEMENT THESE RECEIVER FUNCTIONS ONLY ONCE
// func (p *Pet) says() string {
// 	return p.Sound
// }

// func (p *Pet) legs() int {
// 	return p.NumLegs
// }

// NEED TO DO SEPARATE IMPLEMENTATIONS FOR EACH TYPE
func (d *Dog) says() string {
	return d.Sound
}

func (d *Dog) legs() int {
	return d.NumLegs
}

func (c *Cat) says() string {
	return c.Sound
}

func (c *Cat) legs() int {
	return c.NumLegs
}


func riddle(d Dog) {
	riddle := fmt.Sprintf(`DIRECT: This animal says "%s" and has %d legs. What animal is it?`, d.Sound, d.NumLegs)
	fmt.Println(riddle)
}

func interfacedRiddle(p Pet) {
	riddle := fmt.Sprintf(`VIA I'FACE: This animal says "%s" and has %d legs. What animal is it?`, p.says(), p.legs())
	fmt.Println(riddle)
}

func Interfaces() {
	util.Section("Interfaces:")

	dog := Dog{
		Name: "dog",
		Sound: "woof",
		NumLegs: 4,
	}
	riddle(dog)

	cat := Cat{
		Name: "cat",
		Sound: "meow",
		NumLegs: 4,
		HasTail: true,
	}

	interfacedRiddle(&dog)
	interfacedRiddle(&cat)
}
