package gamenumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameNumber(t *testing.T) {
	assert.Equal(t, "1", NewGameNumber(1).String())
	assert.Equal(t, "Fizz", NewGameNumber(3).String())
	assert.Equal(t, "Buzz", NewGameNumber(5).String())
	assert.Equal(t, "FizzBuzz", NewGameNumber(15).String())
}
