package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 需求：
// 如果一个数能被 3 整除，那么这个数字就是 "Fizz"
// 如果一个数能被 5 整除，那么这个数字就是 "Buzz"
// 如果一个数既能被 3 整除，又能被 5 整除，那么这个数字就是 "FizzBuzz"
func TestFizzbuzz(t *testing.T) {
	assert.Equal(t, "1", Fizzbuzz(1))
	assert.Equal(t, "Fizz", Fizzbuzz(3))
	assert.Equal(t, "Buzz", Fizzbuzz(5))
	assert.Equal(t, "FizzBuzz", Fizzbuzz(15))
}
