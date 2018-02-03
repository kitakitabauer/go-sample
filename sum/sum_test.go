package stdout

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct{
		name string
		x int
		y int
		out int
	}{
		{
			name: "2 + 3 = 5",
			x: 2,
			y: 3,
			out: 5,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			out := sum(v.x, v.y)
			if out != v.out {
				t.Errorf("input(%d, %d) = %d, want %d", v.x, v.y, out, v.out)
			}
		})
	}
}
