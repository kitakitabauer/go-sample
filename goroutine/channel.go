package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ { // 2つのgoroutine(ワーカー)を生成
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "http://www.example.com"
	queue <- "http://www.example.net"
	queue <- "http://www.example.net/foo"
	queue <- "http://www.example.net/bar"

	close(queue) // goroutineに終了を伝える
	wg.Wait() // 全てのgoroutineが終了するのを待つ
}

func fetchURL(queue chan string) {
	for {
		url, more := <- queue // closeされると more が falseになる
		if more {
			// url取得処理
			fmt.Println("fetching", url)
		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}
	}
}