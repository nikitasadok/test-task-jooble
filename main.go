package main

import (
	"calculator/calculator"
	"fmt"
)

func main() {
	c := calculator.NewCalculator("- -")

	fmt.Println(c.Evaluate())
}
