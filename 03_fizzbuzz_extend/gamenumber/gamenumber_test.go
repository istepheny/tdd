package gamenumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 需求：
// 如果一个数能被 3 整除，或者包含数字 3，那么这个数字就是 "Fizz"
// 如果一个数能被 5 整除，或者包含数字 5，那么这个数字就是 "Buzz"
// 如果一个数能被 3 整除，或者包含数字 5，那么这个数字就是 "FizzBuzz"
// 如果一个数能被 5 整除，或者包含数字 3，那么这个数字就是 "FizzBuzz"
func TestGameNumber(t *testing.T) {
	assert.Equal(t, "1", NewGameNumber(1).String())
	assert.Equal(t, "Fizz", NewGameNumber(3).String())
	assert.Equal(t, "Fizz", NewGameNumber(13).String())
	assert.Equal(t, "FizzBuzz", NewGameNumber(51).String())
	assert.Equal(t, "Buzz", NewGameNumber(52).String())
	assert.Equal(t, "FizzBuzz", NewGameNumber(53).String())
}
