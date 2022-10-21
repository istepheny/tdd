package length

import (
	"github.com/stretchr/testify/assert"
	"tdd/05_unit_converter_refactor/unit"
	"testing"
)

// 需求：
// 单位转换
// 1 foot = 12 inch
// 1 yard = 36 inch
func TestLength(t *testing.T) {
	assert.Equal(t, NewLength(1, unit.Inch).String(), "1 (Inch)")
	assert.Equal(t, NewLength(1, unit.Foot).String(), "1 (Foot)")
	assert.Equal(t, NewLength(1, unit.Yard).String(), "1 (Yard)")
	assert.Equal(t, true, NewLength(1, unit.Foot).Equal(NewLength(12, unit.Inch)))
	assert.Equal(t, false, NewLength(1, unit.Foot).Equal(NewLength(13, unit.Inch)))
	assert.Equal(t, true, NewLength(1, unit.Yard).Equal(NewLength(36, unit.Inch)))
	assert.Equal(t, false, NewLength(1, unit.Yard).Equal(NewLength(37, unit.Inch)))
}
