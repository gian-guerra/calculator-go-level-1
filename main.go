package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gian-guerra/calculator-go-level-1/calculator"
)

func main() {
	calc := calculator.NewCalculator()

	calc.Register("add", calculator.Add{})

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <operation> <num1> <num2>")
		return
	}

	opName := os.Args[1]

	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("First number invalid: ", err)
		return
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Second number invalid", err)
		return
	}

	result, err := calc.Execute(opName, a, b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}
