package game

import "tdd/02_fizzbuzz_refactor/gamenumber"

type Game struct {
	gameNumbers []*gamenumber.GameNumber
}

func NewGame(size int) *Game {
	gameNumbers := []*gamenumber.GameNumber{}
	for i := 1; i <= size; i++ {
		gameNumber := gamenumber.NewGameNumber(i)
		gameNumbers = append(gameNumbers, gameNumber)
	}

	return &Game{gameNumbers: gameNumbers}
}

func (g *Game) Size() int {
	return len(g.gameNumbers)
}

func (g *Game) Words() (words []string) {
	for _, gameNumber := range g.gameNumbers {
		words = append(words, gameNumber.String())
	}

	return words
}
