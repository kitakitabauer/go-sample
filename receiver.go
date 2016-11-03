package main

import (
	. "fmt"
)

// Go 言語では構造体は値
type Person struct{
	Name string
}

// Person 型に対してメソッドを定義
// Person という値型に対してメソッド定義されたものを"値レシーバ"と呼ぶ
func (p Person) Greet(msg string) {
	Printf("%s, I'm %s.\n", msg, p.Name)
}

// *Person 型に対してメソッドを定義
// *Person というポインタ型に対してメソッド定義されたものを"ポインタレシーバ"と呼ぶ
func (pp *Person) Shout(msg string) {
	Printf("%s!!!\n", msg)
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

// p の値を変更したいのでポインタレシーバで定義
func (p *Person) ChangeNameUnsafe(name string) {
	// これだけだと危険
	p.Name = name
}

func (p *Person) ChangeNameSafe(name string) {
	// 他の言語だと this == nil みたいなものなので違和感はあるが、
	// ガード句を設けておけば安全
	// これがないと実行時エラーでpanicになる危険性がある
	if p == nil {
		return
	}
	p.Name = name
}

func main() {
	p := Person{Name: "Taro"} // 値型の変数を準備
	p.Greet("Hi") // => Hi, I'm Taro.

	pp := &Person{Name: "Taro"} // ポインタ型の変数を準備
	pp.Shout("Oh My God") // => Oh My God!!!

	// コンパイラによるレシーバの暗黙的変換
	(*pp).Greet("Hello") // => Hello, I'm Taro. "*pp"は"p"という値を参照するため、値型のGreet()メソッド呼び出しが可能
	pp.Greet("Hello") // => Hi, I'm Taro. コンパイラが上のコードに変換してくれる

	(&p).Shout("Oh My Gosh") // => Oh My Gosh!!! "&p"は"*Person"のアドレスを参照するため、ポインタ型のShout()メソッド呼び出しが可能
	p.Shout("Oh My God") // => Oh My God!!! コンパイラが上のコードに変換してくれる

	// ポインタレシーバのメソッドは、nil ポインタ変数からでも呼び出しが可能!フィールド参照がある場合はNG
	var nilp *Person
	nilp.Shout("Fantastic") // => Fantastic!!!

	pp = NewPerson("Taro")
	pp.ChangeNameUnsafe("Jiro")

	pp.Greet("Hi") // Hi, I'm Jiro. ChangeName()メソッドのレシーバは *Person なので、 p.nameフィールドを書き換えることができる。

	// nilp.ChangeNameUnsafe("Jiro") // => panic: runtime error: invalid memory address or nil pointer deference
	nilp.ChangeNameSafe("Jiro") // 安全
}


