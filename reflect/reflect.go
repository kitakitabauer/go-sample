package reflect

import (
	"reflect"
	"fmt"
)

func call(f interface{}) {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		panic("f must be func.")
	}
	fv.Call([]reflect.Value{})
}

func loop(fs ...interface{}) {
	for _, v := range fs {
		f := func() {
			call(v)
		}
		f()
	}
}

func main() {
	f1 := func () {
		fmt.Println("Hello")
	}
	f2 := func () {
		fmt.Println("World!!")
	}

	loop(f1, f2)
}
