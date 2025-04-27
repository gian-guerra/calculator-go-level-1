package calculator

type Operation interface {
	Execute(a, b float64) float64
	Name() string
}

type Add struct{}

func (Add) Execute(a, b float64) float64 {
	return a + b
}

func (Add) Name() string {
	return "Addition"
}

type Subtract struct{}

func (Subtract) Execute(a, b float64) float64 {
	return a - b
}

func (Subtract) Name() string {
	return "Subtraction"
}

type Multiply struct{}

func (Multiply) Execute(a, b float64) float64 {
	return a * b
}

func (Multiply) Name() string {
	return "Multiplication"
}
