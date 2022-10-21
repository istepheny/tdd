package gamenumber

import (
	"strconv"
	"strings"
)

type GameNumber struct {
	rawNumber int
}

func NewGameNumber(i int) *GameNumber {
	return &GameNumber{rawNumber: i}
}

func (g *GameNumber) String() string {
	if g.isRelatedTo(3) && g.isRelatedTo(5) {
		return "FizzBuzz"
	}

	if g.isRelatedTo(3) {
		return "Fizz"
	}

	if g.isRelatedTo(5) {
		return "Buzz"
	}

	return strconv.Itoa(g.rawNumber)
}

func (g *GameNumber) isRelatedTo(i int) bool {
	return g.rawNumber%i == 0 || strings.Contains(strconv.Itoa(g.rawNumber), strconv.Itoa(i))
}
