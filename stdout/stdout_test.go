package main

import (
	"bytes"
	"testing"
)

var buf *bytes.Buffer

func init() {
	buf = &bytes.Buffer{}
}

func TestFprint(t *testing.T) {
	fprint(buf)
	output := buf.String()
	correct := "writer\n"

	if correct != output {
		t.Errorf("got %s, but want: %s", correct, output)
		t.FailNow()
	}
}
