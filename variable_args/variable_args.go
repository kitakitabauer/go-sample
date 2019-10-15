package variable_args

import "fmt"

type Struct struct {
	ID string
}

func main() {
	// d := []Struct{
	// 	{
	// 		ID: "string",
	// 	},
	// 	{
	// 		ID: "string",
	// 	},
	// }
	d := make([]Struct, 0)
	d = append(d, Struct{ID: "string"})
	d = append(d, Struct{ID: "string"})
	hoge(d)
	return
}

func hoge(d []Struct) {
	fuga(d...)
	return
}

func fuga(d ...Struct) {
	fmt.Println(d)
	return
}
