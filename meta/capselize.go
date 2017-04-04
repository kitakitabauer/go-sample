package meta

import "fmt"

type Human struct {
	name string // package scope
}

// NewHuman is constructor
func NewHuman(name string) *Human {
	return &Human{ name: name }
}

func (h Human) GetName() string {
	return h.name
}

func (h *Human) SetName(name string) {
	h.name = name
}

func Humanize() {
	human := NewHuman("alice")
	fmt.Println(human.GetName())

	human.SetName("bob")
	fmt.Println(human.GetName())
}