package day1

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func compute(f io.Reader, decode func([]string) (move, move)) (int, error) {
	ir, err := newScoreReader(f)
	if err != nil {
		return 0, err
	}
	var score int
	for {
		n, ok, err := ir.next()
		if err != nil {
			return 0, err
		}
		if !ok {
			break
		}
		mine, theirs := decode(n)
		score += computeScore(mine, theirs)
	}
	return score, nil
}

func decode1(tokens []string) (mine, theirs move) {
	mine = letterToMove[tokens[1]]
	theirs = letterToMove[tokens[0]]
	return
}

func decode2(tokens []string) (mine, theirs move) {
	theirs = letterToMove[tokens[0]]
	res := letterToOutcome[tokens[1]]
	mine = move((int(theirs) + int(res) + 2) % 3)
	return
}

// scoreReader reads the input and returns the sum of consecutive numbers.
type scoreReader struct {
	scanner *bufio.Scanner
}

func newScoreReader(r io.Reader) (*scoreReader, error) {
	return &scoreReader{
		scanner: bufio.NewScanner(r),
	}, nil
}

func (sc *scoreReader) next() ([]string, bool, error) {
	ok := sc.scanner.Scan()
	if !ok {
		return nil, false, sc.scanner.Err()
	}
	l := sc.scanner.Text()
	outs := strings.SplitN(l, " ", 2)
	if len(outs) != 2 {
		return nil, false, fmt.Errorf("invalid entry: %s", l)
	}
	return outs, true, nil
}

func computeScore(mine, theirs move) int {
	return int(3*((mine-theirs+4)%3) + mine + 1)
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
