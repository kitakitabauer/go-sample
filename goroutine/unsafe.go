package main

import (
	"fmt"
	"time"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// 問題：同時に複数のgoroutineからnにアクセス
	// → nの値がバラバラ
	// 解決方法：1つのgoroutineからアクセス
	// channelを使う
	n := 1
	go func() {
		for i := 2; i <= 5; i++ {
			fmt.Println(n, "*", i)
			n *= i
			time.Sleep(100)
		}
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println(n, "+", i)
		n += 1
		time.Sleep(100)
	}
}
