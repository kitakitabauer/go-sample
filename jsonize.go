package main

import (
	"encoding/json"
	"fmt"
)

type M map[string]interface{}
type A []interface{}

func main() {
	values := M{
		"a": "a",
		"b": func() A {
			return A{1, 2, 3}
		}(),
		"c": A{"va", "vb", "vc"},
		"d": M{
			"dd": 1.1,
			"cc": 2.2,
			"ee": nil,
			"qq": true}}

	bytes, err := json.MarshalIndent(values, "", "    ")
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	jsonstring := string(bytes)
	fmt.Println(jsonstring)

	parse()
}

type Hoge struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

func parse() {
	str := `{"code": "202", "error": "hoge"}`
	hoge := struct {
		Code  string `json:"code"`
		Error string `json:"error"`
	}{}
	//var hoge Hoge
	err := json.Unmarshal([]byte(str), &hoge)
	fmt.Printf("%#v\n", err)
	fmt.Printf("%#v", hoge)
	return
}
