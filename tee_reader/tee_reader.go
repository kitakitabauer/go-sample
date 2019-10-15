package tee_reader

import (
	"os"
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	// open files r and w
	r, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	w, err := os.Create("tee_reader_output.txt")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// we want to copy the file
	// and at the same time create a MD5 checksum

	hash := md5.New()

	n, err := io.Copy(w, io.TeeReader(r, hash))
	if err != nil {
		panic(err)
	}

	// the end result ?
	// kill 2 birds with 1 stone thanks to io.TeeReader()
	fmt.Printf("Copied %v bytes with checksum %x \n", n, hash.Sum(nil))
}
