package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnit(t *testing.T) {
	assert.Equal(t, "Inch", Inch.Name())
	assert.Equal(t, 1, Inch.TransferRateToInch())
	assert.Equal(t, "Foot", Foot.Name())
	assert.Equal(t, 12, Foot.TransferRateToInch())
	assert.Equal(t, "Yard", Yard.Name())
	assert.Equal(t, 36, Yard.TransferRateToInch())
}
