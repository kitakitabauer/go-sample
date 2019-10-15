package _struct

import (
	. "fmt"
	"reflect"
)

func main() {
	hoge := Hoge{100, "hello"}
	m := StructToMap(&hoge)

	Println(m)

}

type Hoge struct {
	Fuga int
	Huga string
}

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}
