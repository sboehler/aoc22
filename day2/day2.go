package day1

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func compute(f io.Reader) (int, error) {
	ir, err := newInputReader(f)
	if err != nil {
		return 0, err
	}
	var s scorer
	for {
		n, ok, err := ir.next1()
		if err != nil {
			return 0, err
		}
		if !ok {
			break
		}
		s.update1(n)
	}
	return s.score, nil
}

func compute2(f io.Reader) (int, error) {
	ir, err := newInputReader(f)
	if err != nil {
		return 0, err
	}
	var s scorer
	for {
		n, ok, err := ir.next2()
		if err != nil {
			return 0, err
		}
		if !ok {
			break
		}
		s.update2(n)
	}
	return s.score, nil
}

// max keeps track of the sum of the maximum len(max) numbers.
type scorer struct {
	score int
}

func (s *scorer) update1(n round1) {
	s.score += score(n)
}

func (s *scorer) update2(n round2) {
	mine := move((int(n.theirs) + int(n.res) - 1) % 3)
	if mine < 0 {
		mine += 3
	}
	s.score += score(round1{
		mine:   mine,
		theirs: n.theirs,
	})
}

func score(o round1) int {
	ms := int(o.mine) + 1
	switch (o.mine - o.theirs) % 3 {
	case 0:
		return 3 + ms
	case 1, -2:
		return 6 + ms
	case 2, -1:
		return ms
	}
	panic("impossible")
}

type move int

const (
	Rock move = iota
	Paper
	Scissors
)

type outcome int

const (
	Lose outcome = iota
	Draw
	Win
)

var letterToMove = map[string]move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var letterToOutcome = map[string]outcome{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

type round1 struct{ mine, theirs move }

type round2 struct {
	theirs move
	res    outcome
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

func (ir *inputReader) next() ([]string, bool, error) {
	ok := ir.scanner.Scan()
	if !ok {
		return nil, false, ir.scanner.Err()
	}
	l := ir.scanner.Text()
	outs := strings.SplitN(l, " ", 2)
	if len(outs) != 2 {
		return nil, false, fmt.Errorf("invalid entry: %s", l)
	}
	return outs, true, nil
}

func (ir *inputReader) next1() (round1, bool, error) {
	outs, ok, err := ir.next()
	if !ok || err != nil {
		return round1{}, ok, err
	}
	return round1{
		theirs: letterToMove[outs[0]],
		mine:   letterToMove[outs[1]],
	}, true, nil
}

func (ir *inputReader) next2() (round2, bool, error) {
	outs, ok, err := ir.next()
	if !ok || err != nil {
		return round2{}, ok, err
	}
	return round2{
		theirs: letterToMove[outs[0]],
		res:    letterToOutcome[outs[1]],
	}, true, nil
}
