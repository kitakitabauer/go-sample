package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"strings"
	"log"
	"unicode"
	"strconv"
)

var (
	dict map[string]struct{}
	sum int
)

// 標準入力で受け取ったスペース区切りの単語において
// 同じ文字は除き、先頭が大文字 または 数値の単語の数を数えて標準出力する。
func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text :=  sc.Text()

		reset()
		var checkWords []string
		allWords := strings.Fields(text)
		for _, word := range allWords {
			if _, ok := dict[word]; !ok {
				dict[word] = struct{}{}
				checkWords = append(checkWords, word)
			}
		}

		for _, word := range checkWords {
			if isFirstUpper(word) {
				sum++
			} else if checkType(word) {
				sum++
			}
		}
		fprint(os.Stdout, strconv.Itoa(sum))

		dict = make(map[string]struct{})
		sum = 0
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}
}

func isFirstUpper(w string) bool {
	if w == "" {
		return false;
	}
	r := rune(w[0])
	return unicode.IsUpper(r)
}

func checkType(w string) bool {
	_, err := strconv.Atoi(w)
	if err != nil {
		return false
	}
	return true
}

func fprint(w io.Writer, ans string) {
	fmt.Fprint(w, ans)
}

func reset() {
	dict = make(map[string]struct{})
	sum = 0
}
