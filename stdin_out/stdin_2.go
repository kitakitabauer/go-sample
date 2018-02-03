package main

import (
	"bufio"
	"os"
	"strings"
	"log"
)

// 標準入力で受け取ったカンマ区切りの数値において、
// 奇数番目の要素を除いて出力する。ただし前の数値と同じものは除かない。
func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text :=  sc.Text()

		var answer string
		var counter int
		var prev string
		nums := strings.Split(text, ",")

		for _, num := range nums {
			if counter % 2 == 0 {
				// odd number
				if prev == num {
					answer = addComma(answer)
					answer += num
				}
			} else {
				// even number
				answer = addComma(answer)
				answer += num
			}
			prev = num
			counter++
		}

		fprint(os.Stdout, answer)
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}
}

//func fprint(w io.Writer, ans string) {
//	fmt.Fprint(w, ans)
//}

func addComma(ans string) string {
	if len(ans) > 0 {
		ans = ans + ","
	}
	return ans
}
