package day3

import (
	"io"
	"os"
	"testing"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		compute func(r io.Reader) (int, error)
		want    int
	}{
		{
			name:    "Part 1 (example)",
			file:    "example.txt",
			compute: compute,
			want:    157,
		},
		{
			name:    "Part 1 (input)",
			file:    "input.txt",
			compute: compute,
			want:    7831,
		},
		{
			name:    "Part 2 (example)",
			file:    "example.txt",
			compute: compute2,
			want:    70,
		},
		{
			name:    "Part 2 (input)",
			file:    "input.txt",
			compute: compute2,
			want:    2683,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			got, err := test.compute(f)

			if err != nil {
				t.Fatalf("compute() returned error %v,  want nil", err)
			}
			if got != test.want {
				t.Errorf("compute() = %d, want %d", got, test.want)
			}
		})
	}
}
