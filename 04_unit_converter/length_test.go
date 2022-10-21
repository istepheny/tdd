package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 需求：
// 单位转换
// 1 foot = 12 inch
// 1 yard = 36 inch
func TestLength(t *testing.T) {
	assert.Equal(t, NewLength(1, "Inch").String(), "1 (Inch)")
	assert.Equal(t, NewLength(1, "Foot").String(), "1 (Foot)")
	assert.Equal(t, NewLength(1, "Yard").String(), "1 (Yard)")
	assert.Equal(t, true, NewLength(1, "Foot").Equal(NewLength(12, "Inch")))
	assert.Equal(t, false, NewLength(1, "Foot").Equal(NewLength(13, "Inch")))
	assert.Equal(t, true, NewLength(1, "Yard").Equal(NewLength(36, "Inch")))
	assert.Equal(t, false, NewLength(1, "Yard").Equal(NewLength(37, "Inch")))
}
