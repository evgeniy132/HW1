package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name        string
	Age         int
	Height      int
	Hobby       []string
	Description string
	Gender      Gender
}

type Gender string

const (
	male   Gender = "man"
	female Gender = "woman"
)

func (g Gender) String() string {
	return string(g)
}

func NewPerson(name, description string, gender Gender, age, height int, Hobbies ...string) Person {
	return Person{Name: name, Age: age, Height: height, Description: description, Hobby: Hobbies, Gender: gender}
}
func (p Person) TellBio() {
	fmt.Printf(
		" Hello, my name is %s,\n I`am %d years old\n And my hobbies are: %s\n "+
			"I can tell u about me: %s\n My gender is %s  ",
		p.Name, p.Age, strings.Join(p.Hobby, " and "), p.Description, p.Gender)
}

func main() {
	var p1 = NewPerson("Tolia", "Опис до дескріпшина", male, 22, 178, "sport", "fishing", "programming")

	p1.TellBio()
}
