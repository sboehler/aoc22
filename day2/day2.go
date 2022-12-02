package day1

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func compute(f io.Reader, n int) (int, error) {
	ir, err := newInputReader(f)
	if err != nil {
		return 0, err
	}
	var s scorer
	for {
		n, ok, err := ir.next()
		if err != nil {
			return 0, err
		}
		if !ok {
			break
		}
		s.update(n)
		fmt.Println(n)
	}
	return s.score, nil
}

// max keeps track of the sum of the maximum len(max) numbers.
type scorer struct {
	score int
}

func (m *scorer) update(n outcome) {
	m.score += score(n)
}

func score(o outcome) int {
	switch (o.mine - o.theirs) % 3 {
	case 0:
		return 3 + (o.mine + 1)
	case 1, -2:
		return 6 + (o.mine + 1)
	case 2, -1:
		return o.mine + 1
	}
	panic("impossible")
}

const (
	Rock int = iota
	Paper
	Scissors
)

type outcome struct{ mine, theirs int }

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

func (ir *inputReader) next() (outcome, bool, error) {
	ok := ir.scanner.Scan()
	if !ok {
		return outcome{}, false, ir.scanner.Err()
	}
	l := ir.scanner.Text()
	outs := strings.SplitN(l, " ", 2)
	if len(outs) != 2 {
		return outcome{}, false, fmt.Errorf("invalid entry: %s", l)
	}
	var mine, theirs int
	switch outs[0] {
	case "A":
		theirs = Rock
	case "B":
		theirs = Paper
	case "C":
		theirs = Scissors
	default:
		return outcome{}, false, fmt.Errorf("invalid move: %s", outs[0])
	}
	switch outs[1] {
	case "X":
		mine = Rock
	case "Y":
		mine = Paper
	case "Z":
		mine = Scissors
	default:
		return outcome{}, false, fmt.Errorf("invalid move: %s", outs[1])
	}
	return outcome{mine, theirs}, true, nil
}
