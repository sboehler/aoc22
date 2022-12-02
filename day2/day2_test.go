package day1

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name   string
		file   string
		decode func([]string) (move, move)
		want   int
	}{
		{
			name:   "Part 1 (example)",
			file:   "example.txt",
			decode: decode1,
			want:   15,
		},
		{
			name:   "Part 1 (solution)",
			file:   "input.txt",
			decode: decode1,
			want:   8392,
		},
		{
			name:   "Part 2 (example)",
			file:   "example.txt",
			decode: decode2,
			want:   12,
		},
		{
			name:   "Part 2 (solution)",
			file:   "input.txt",
			decode: decode2,
			want:   10116,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			got, err := compute(f, test.decode)

			if err != nil {
				t.Fatalf("compute() returned error %v,  want nil", err)
			}
			if got != test.want {
				t.Errorf("compute() = %d, want %d", got, test.want)
			}
		})
	}
}
