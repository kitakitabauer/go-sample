package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	fprint(os.Stdout)
}

func fprint(w io.Writer) {
	fmt.Fprint(w, "writer\n")
}
