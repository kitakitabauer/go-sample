package meta

import "fmt"

type Speaker interface {
	Speak()
}

func (h Human) Speak() {
	fmt.Println("I'm a perfect Human.")
}

// Woman not embed
type Woman struct {
	name string
}

func NewWoman(name string) *Woman {
	return &Woman{name: name}
}

func (w Woman) Speak() {
	fmt.Println("I'm a perfect Woman.")
}

func HumanizeByPolymorphism() {
	human := NewHuman("alice")
	woman := NewWoman("emily")

	speak(human)
	speak(woman)
}

func speak(speaker Speaker) {
	speaker.Speak()
}