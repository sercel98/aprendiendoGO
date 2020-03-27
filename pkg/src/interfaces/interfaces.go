package main

import "fmt"

type animal interface {
	speak() string
}

type dog struct {
	name  string
	sound string
}

type cat struct {
	name  string
	sound string
}

type cow struct {
	name  string
	sound string
}

//Here is where dog, cat and cow implement the interface method. So, now they are all animals.
func (d dog) speak() string {
	return d.sound + ", my name is " + d.name + " and I'm a dog!"
}

func (c cat) speak() string {
	return c.sound + ", my name is " + c.name + " and I'm a cat!"
}

func (c cow) speak() string {
	return c.sound + ", my name is " + c.name + " and I'm a cow!"
}

func speakAnimal(a animal) {
	fmt.Println(a.speak())
}
func main() {
	aDog := dog{
		name:  "Toby",
		sound: "Guau",
	}
	aCat := cat{
		name:  "Michi",
		sound: "Miau",
	}
	aCow := cow{
		name:  "Milu",
		sound: "Muuu",
	}
	speakAnimal(aDog)
	speakAnimal(aCow)
	speakAnimal(aCat)

}
