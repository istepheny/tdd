package length

import (
	"fmt"
	"tdd/05_unit_converter_refactor/unit"
)

type Length struct {
	amount int
	unit   unit.Unit
}

func NewLength(amount int, unit unit.Unit) *Length {
	return &Length{
		amount: amount,
		unit:   unit,
	}
}

func (l *Length) Equal(another *Length) bool {
	return l.getAmountInInch() == another.getAmountInInch()
}

func (l *Length) getAmountInInch() int {
	return l.amount * l.unit.TransferRateToInch()
}

func (l *Length) String() string {
	return fmt.Sprintf("%d (%s)", l.amount, l.unit.Name())
}
