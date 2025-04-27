package calculator

import (
	"fmt"
)

type Calculator struct {
	operations map[string]Operation
}

func NewCalculator() *Calculator {
	return &Calculator{
		operations: make(map[string]Operation),
	}
}

func (c *Calculator) Register(name string, op Operation) {
	c.operations[name] = op
}

func (c *Calculator) Execute(name string, a, b float64) (float64, error) {
	op, exists := c.operations[name]
	if !exists {
		return 0, fmt.Errorf("operation '%s' not found", name)
	}

	return op.Execute(a, b), nil
}
