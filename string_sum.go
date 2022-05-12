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
	if input == "" {
		return "", makeError(errorEmptyInput)
	}

	stringA := ""
	var operator rune
	for i, s := range input[1:] {
		if s != '-' && s != '+' {
			continue
		}

		if stringA != "" {
			return "", makeError(errorNotTwoOperands)
		}

		stringA = strings.TrimSpace(input[:i+1])
		operator = s
	}

	a, err := strconv.Atoi(stringA)
	if err != nil {
		return "", makeError(err)
	}

	stringB := strings.TrimSpace(input[len(stringA)+1:])
	b, err := strconv.Atoi(stringB)
	if err != nil {
		return "", makeError(err)
	}

	var result int
	switch operator {
	case '-':
		result = a - b
	case '+':
		result = a + b
	}

	return strconv.Itoa(result), nil
}

func makeError(err error) error {
	return fmt.Errorf("cannot calculate sum: %w", err)
}
