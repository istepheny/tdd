package unit

var (
	Inch = &inch{}
	Foot = &foot{}
	Yard = &yard{}
)

type Unit interface {
	Name() string
	TransferRateToInch() int
}

type inch struct{}

func (inch) Name() string {
	return "Inch"
}

func (inch) TransferRateToInch() int {
	return 1
}

type foot struct{}

func (foot) Name() string {
	return "Foot"
}

func (foot) TransferRateToInch() int {
	return 12
}

type yard struct{}

func (yard) Name() string {
	return "Yard"
}

func (yard) TransferRateToInch() int {
	return 36
}
