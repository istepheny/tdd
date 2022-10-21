package main

import (
	"fmt"
)

type Length struct {
	amount       int
	amountInInch int
	unit         string
}

func NewLength(amount int, unit string) *Length {
	l := &Length{
		amount: amount,
		unit:   unit,
	}

	if unit == "Foot" {
		l.amountInInch = amount * 12
	} else if unit == "Yard" {
		l.amountInInch = amount * 36
	} else if unit == "Inch" {
		l.amountInInch = amount
	}

	return l
}

func (l *Length) Equal(another *Length) bool {
	return l.amountInInch == another.amountInInch
}

func (l *Length) String() string {
	return fmt.Sprintf("%d (%s)", l.amount, l.unit)
}
