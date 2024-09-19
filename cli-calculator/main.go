package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Perform the calculation based on operator and return the result or error
func calculate(num1 float64, operator string, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("unsupported operator. Use +, -, *, or /")
	}
}

func main() {
	// Handle command-line input
	if len(os.Args) != 4 {
		fmt.Println("Usage: <number1> <operator> <number2>")
		return
	}

	// Parse the inputs
	num1Str := os.Args[1]
	operator := os.Args[2]
	num2Str := os.Args[3]

	// Convert string to float
	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	// Check if inputs are valid numbers
	if err1 != nil || err2 != nil {
		fmt.Println("Please enter valid numbers.")
		return
	}

	// Call the calculate function and handle the result
	result, err := calculate(num1, operator, num2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
