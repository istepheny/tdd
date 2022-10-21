package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	game := NewGame(100)
	assert.Equal(t, 100, game.Size())

	wordsToBeSpoken := game.Words()
	assert.Equal(t, 100, len(wordsToBeSpoken))
	assert.Equal(t, "1", wordsToBeSpoken[0])
	assert.Equal(t, "Fizz", wordsToBeSpoken[2])
	assert.Equal(t, "Buzz", wordsToBeSpoken[4])
	assert.Equal(t, "FizzBuzz", wordsToBeSpoken[14])
}
