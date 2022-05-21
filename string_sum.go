package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	inputArray := strings.Fields(input)
	input = strings.Join(inputArray, "")
	if len(input) == 0 {
		return "", fmt.Errorf("my error: %w", errorEmptyInput)
	}
	input = strings.TrimLeft(input, "+")
	inputArray = strings.Split(input, "+")
	fmt.Println(inputArray, len(inputArray))

	if len(inputArray) > 2 {
		return "", fmt.Errorf("my error: %w", errorNotTwoOperands)
	}

	if len(inputArray) == 2 {
		num1, err1 := strconv.Atoi(inputArray[0])
		if err1 != nil {
			return "", fmt.Errorf("my error: %w", err1)
		}

		num2, err2 := strconv.Atoi(inputArray[1])
		if err2 != nil {
			return "", fmt.Errorf("my error: %w", err2)
		}

		output = strconv.Itoa(num1 + num2)
		return output, nil
	}

	minus := 1
	if string(input[0]) == "-" {
		minus = -1
		input = strings.TrimLeft(input, "-")
	}
	inputArray = strings.Split(input, "-")

	if len(inputArray) != 2 {
		return "", fmt.Errorf("my error: %w", errorNotTwoOperands)
	}

	num1, err1 := strconv.Atoi(inputArray[0])
	if err1 != nil {
		return "", fmt.Errorf("my error: %w", err1)
	}

	num2, err2 := strconv.Atoi(inputArray[1])
	if err2 != nil {
		return "", fmt.Errorf("my error: %w", err2)
	}

	output = strconv.Itoa(minus*num1 - num2)
	return output, nil
}
