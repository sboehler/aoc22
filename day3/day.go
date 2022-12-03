package day3

import (
	"bufio"
	"fmt"
	"io"
)

func compute(f io.Reader) (int, error) {
	s := bufio.NewScanner(f)
	var res int
	for s.Scan() {
		line := s.Text()
		cmp1, cmp2 := line[:len(line)/2], line[len(line)/2:]
		dups := duplicates(cmp1, cmp2)
		for c := range dups {
			res += score(c)
		}
	}
	return res, nil
}

func compute2(f io.Reader) (int, error) {
	s := bufio.NewScanner(f)
	var res int
	for {
		ls, ok, err := readGroup(s)
		if !ok {
			return res, err
		}
		dups := duplicates(ls[0], ls[1])
		for _, c := range ls[2] {
			if _, ok := dups[c]; ok {
				res += score(c)
				break
			}
		}
	}
}

func readGroup(s *bufio.Scanner) ([]string, bool, error) {
	var res []string
	for i := 0; i < 3; i++ {
		if ok := s.Scan(); !ok {
			return nil, ok, s.Err()
		}
		res = append(res, s.Text())
	}
	return res, true, nil
}

func duplicates(s1, s2 string) map[rune]struct{} {
	cc := make(map[rune]struct{})
	for _, c := range s1 {
		cc[c] = struct{}{}
	}
	res := make(map[rune]struct{})
	for _, c := range s2 {
		if _, ok := cc[c]; ok {
			res[c] = struct{}{}
		}
	}
	return res
}

func score(c rune) int {
	v := int(c)
	switch {
	case 97 <= v && v <= 122:
		return v - 96
	case 65 <= v && v <= 90:
		return v - 38
	default:
		panic(fmt.Sprintf("invalid character: %v", c))
	}
}
