package tsv_reader

import (
	"fmt"
	"os"

	"github.com/dogenzaka/tsv"
)

type Row struct {
	UserID    string
	CreatedAt int
}

func main() {
	file, err := os.Open("tsv_reader_sample.tsv")
	if err != nil {
		error.Error(err)
		return
	}

	row := Row{}
	parser := tsv.NewParserWithoutHeader(file, &row)

	rows := make([]Row, 0)
	for {
		eof, err := parser.Next()
		if eof {
			fmt.Println(len(rows))
			return
		}
		if err != nil {
			error.Error(err)
		}

		fmt.Printf("%#v\n", row)
		rows = append(rows, row)
	}
}
