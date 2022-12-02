package day1

import (
	"bufio"
	"io"
	"strconv"
)

func compute(f io.Reader, n int) (int, error) {
	ir, err := newInputReader(f)
	if err != nil {
		return 0, err
	}
	m := make(max, n)
	for {
		n, ok, err := ir.next()
		if err != nil {
			return 0, err
		}
		if !ok {
			break
		}
		m.update(n)
	}
	return m.sum(), nil
}

// max keeps track of the sum of the maximum len(max) numbers.
type max []int

func (m max) update(n int) {
	for i := range m {
		if n <= m[i] {
			return
		}
		if i > 0 {
			m[i-1] = m[i]
		}
		m[i] = n
	}
}

func (m max) sum() int {
	var res int
	for _, n := range m {
		res += n
	}
	return res
}

// inputReader reads the input and returns the sum of consecutive numbers.
type inputReader struct {
	scanner *bufio.Scanner
}

func newInputReader(r io.Reader) (*inputReader, error) {
	scanner := bufio.NewScanner(r)
	return &inputReader{
		scanner: scanner,
	}, nil
}

func (ir *inputReader) next() (int, bool, error) {
	var sum int
	ok := ir.scanner.Scan()
	if !ok {
		return 0, false, ir.scanner.Err()
	}
	for {
		l := ir.scanner.Text()
		if len(l) == 0 {
			return sum, true, nil
		}
		n, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			return 0, false, err
		}
		sum += int(n)
		ok := ir.scanner.Scan()
		if !ok {
			err := ir.scanner.Err()
			return sum, err == nil, err
		}
	}
}
