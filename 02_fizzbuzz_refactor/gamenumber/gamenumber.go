package gamenumber

import "strconv"

type GameNumber struct {
	rawNumber int
}

func NewGameNumber(i int) *GameNumber {
	return &GameNumber{rawNumber: i}
}

func (g *GameNumber) String() string {
	if g.isDivisibleBy(3) && g.isDivisibleBy(5) {
		return "FizzBuzz"
	}

	if g.isDivisibleBy(3) {
		return "Fizz"
	}

	if g.isDivisibleBy(5) {
		return "Buzz"
	}

	return strconv.Itoa(g.rawNumber)
}

func (g *GameNumber) isDivisibleBy(i int) bool {
	return g.rawNumber%i == 0
}
