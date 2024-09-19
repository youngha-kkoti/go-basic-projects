package main

import (
	"testing"
)

// Test addition
func TestAddition(t *testing.T) {
	result, err := calculate(10, "+", 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 15.0
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test subtraction
func TestSubtraction(t *testing.T) {
	result, err := calculate(10, "-", 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 5.0
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test multiplication
func TestMultiplication(t *testing.T) {
	result, err := calculate(10, "*", 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 50.0
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test division
func TestDivision(t *testing.T) {
	result, err := calculate(10, "/", 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 2.0
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test division by zero
func TestDivisionByZero(t *testing.T) {
	_, err := calculate(10, "/", 0)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
	expectedError := "cannot divide by zero"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

// Test unsupported operator
func TestUnsupportedOperator(t *testing.T) {
	_, err := calculate(10, "^", 5)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
	expectedError := "unsupported operator. Use +, -, *, or /"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

// Test invalid operator (empty string)
func TestEmptyOperator(t *testing.T) {
	_, err := calculate(10, "", 5)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
	expectedError := "unsupported operator. Use +, -, *, or /"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
