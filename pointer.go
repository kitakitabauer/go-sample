package main

import (
	. "fmt"
	"strconv"
)

func inc(i *int) { // int型のポインタを受け取る
	*i++ // 受け取ったポインタ型変数に"*"をつけて使用する。ポインタの参照している先の値を使うことができる。
	Println("inc: i = " + strconv.Itoa(*i))
}

func calc(a int, b *int) {
	// 呼び元には影響しない
	a = a + 1
	// 呼び元にも影響する
	*b = *b + 1
}

func main() {
	num := 10
	Println(num) // 10
	inc(&num) // 呼び出し元は"&"を使ってint型変数のアドレスを渡す。ポインタが値渡しされる。Goは全てポインタ渡し(ポインタのコピー)
	Println(num) // 11


	a, b := 10, 10

	// bはアドレス演算子をつける
	calc(a, &b)

	Println("更新されない:", a)
	Println("更新される:", b)
}
