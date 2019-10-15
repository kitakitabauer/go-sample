package type_assertion

import "fmt"

func main() {
	check("Hello")
}

func check(i interface{}) {
	a, b := i.(string)
	fmt.Println(a) // Hello
	fmt.Println(b) // true
}
