package day1

import (
	"os"
	"testing"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		name string
		file string
		max  int
		want int
	}{
		{
			name: "Part 1 (example)",
			file: "example.txt",
			max:  1,
			want: 24000,
		},
		{
			name: "Part 1 (solution)",
			file: "input.txt",
			max:  1,
			want: 70116,
		},
		{
			name: "Part 2 (example)",
			file: "example.txt",
			max:  3,
			want: 45000,
		},
		{
			name: "Part 2 (solution)",
			file: "input.txt",
			max:  3,
			want: 206582,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			got, err := compute(f, test.max)

			if err != nil {
				t.Fatalf("compute() returned error %v,  want nil", err)
			}
			if got != test.want {
				t.Errorf("compute() = %d, want %d", got, test.want)
			}
		})
	}
}
