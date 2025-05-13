package main

import (
	"fmt"
	"math"
)

// Valid operators for calculations. MUST be in priority order.
var operators = []rune{'^', '*', '/', '+', '-'}

// Check if a rune is present in a string, and returns the index or -1 if it isn't present
func checkRuneMembership(stringToSearch string, runeToFind rune) int {
	for i, currentLetter := range stringToSearch {
		if currentLetter == runeToFind {
			return i
		}
	}
	return -1
}

func simpleOperation(operator rune, a, b float64) (float64, error) {
	switch operator {
	case '^':
		return math.Pow(a, b), nil
	case '/':
		if b == 0 {
			return -1, fmt.Errorf("Can't Divide by 0")
		}
		return a / b, nil
	case '*':
		return a * b, nil
	case '-':
		return a - b, nil
	case '+':
		return a + b, nil
	}
	return 0, fmt.Errorf("Wrong operator")
}

func simplifyOperation(expression string, operatorIndex int) string {
	expressionRunes := []rune(expression)
	if operatorIndex == 1 {
		a := float64(expressionRunes[0] - '0')
		b := float64(expressionRunes[2] - '0')

		expressionEnd := expression[3:]

		result, err := simpleOperation(expressionRunes[1], a, b)
		if err != nil {
			fmt.Errorf("Operation error")
			return "error"
		} else {
			return fmt.Sprintf("%f", result) + expressionEnd
		}
	} else {
		// TODO: custom splitting and calculations
	}
}

func evaluateExpression(expression string) string {
	isCompleted := true
	for _, operator := range operators {
		if i := checkRuneMembership(expression, operator); i != -1 {
			isCompleted = false

		}
	}
	if isCompleted {
		fmt.Printf("= %s", expression)
	} else {
		expression = evaluateExpression(expression)
	}
	return expression
}

func main() {
	expression := "hello, this is a test"
	for _, v := range expression {
		fmt.Printf("%c ", v)
	}
	fmt.Println("")
	// test := []rune("hi, my name is Simon, I divide (/)")
	// for i := 0; i < len(test); i++ {
	// 	fmt.Printf("%c", test[i])
	// }
}
