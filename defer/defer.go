package _defer

import "fmt"

type hoge struct {
	ID string
}

func _defer(h *hoge) {
	if h.ID == "" {
		h.ID = "test!"
	}
	fmt.Printf("_deferの中: %#v", h)
}

func main() {
	h := hoge{}
	defer _defer(&h)

	fmt.Printf("main: %#v", h)
}
