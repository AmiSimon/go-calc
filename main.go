package main

import (
	"fmt"
	"log"
	"math"
	"os"
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

func operations(operator rune, a, b float64) (float64, error) {
	switch operator {
	case '^':
		return math.Pow(a, b), nil
	case '/':
		if b == 0 {
			return -1, fmt.Errorf("can't Divide by 0")
		}
		return a / b, nil
	case '*':
		return a * b, nil
	case '-':
		return a - b, nil
	case '+':
		return a + b, nil
	}
	return 0, fmt.Errorf("wrong operator")
}

func simplifyOperation(expression string, operatorIndex int) string {
	expressionEnd := expression[operatorIndex+2:]
	if operatorIndex == 1 {
		return singleStringOperation([]rune(expression), operatorIndex) + expressionEnd
	} else {
		expressionStart := expression[:(operatorIndex - 1)]
		return expressionStart + singleStringOperation([]rune(expression), operatorIndex) + expressionEnd
	}
}

func singleStringOperation(expression []rune, operatorIndex int) string {
	a := float64(expression[operatorIndex-1] - '0')
	b := float64(expression[operatorIndex+1] - '0')

	result, err := operations(expression[operatorIndex], a, b)

	//fmt.Print("a: %f, b: %f, operan")
	if err != nil {
		log.Fatal(err)
		return "error"
	} else {
		return fmt.Sprintf("%.0f", result)
	}
}

func evaluateExpression(expression string) string {
	isCompleted := true
	for _, operator := range operators {
		if i := checkRuneMembership(expression, operator); i != -1 {
			expression = simplifyOperation(expression, i)
		}
		if i := checkRuneMembership(expression, operator); i != -1 {
			isCompleted = false
		}
	}
	if isCompleted {
		fmt.Printf("= %s\n", expression)
	} else {
		expression = evaluateExpression(expression)
	}
	return expression
}

func main() {
	if len(os.Args) > 2 {
		log.Fatal("only one argument must be used :\n    calc [OPERATION]")
		return
	}
	if len(os.Args) == 1 {
		log.Fatal("Not enough arguments for calc, try using :\n    calc [OPERATION]")
		return
	}
	inputExpression := os.Args[1]

	evaluateExpression(inputExpression)
}
