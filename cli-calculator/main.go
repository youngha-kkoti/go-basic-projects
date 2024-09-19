package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args[0] is the program name, so actual arguments start from 1
	if len(os.Args) != 4 {
		fmt.Println("Usage: <number1> <operator> <number2>")
		return
	}

	// Parsing command-line arguments
	num1Str := os.Args[1]
	operator := os.Args[2]
	num2Str := os.Args[3]

	// Convert strings to numbers
	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	// Check if input values are valid numbers
	if err1 != nil || err2 != nil {
		fmt.Println("Please enter valid numbers.")
		return
	}

	// Perform calculation
	calculate(num1, operator, num2)
}

func calculate(num1 float64, operator string, num2 float64) {
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero.")
		} else {
			fmt.Printf("Result: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Unsupported operator. Use +, -, *, or /.")
	}
}
