package calculator

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidExpression = errors.New("invalid expression")

func Calculate(expression string) (string, error) {
	if len(expression) == 0 {
		return "", ErrInvalidExpression
	}

	for _, ch := range expression {
		if !unicode.IsDigit(ch) && ch != '+' && ch != '-' && ch != '*' && ch != '/' {
			return "", ErrInvalidExpression
		}
	}

	result := 0
	currentNum := 0
	currentOp := '+'

	for i, ch := range expression {
		if unicode.IsDigit(ch) {
			digit, _ := strconv.Atoi(string(ch))
			currentNum = currentNum*10 + digit
		}

		if (!unicode.IsDigit(ch) && ch != ' ') || i == len(expression)-1 {
			switch currentOp {
			case '+':
				result += currentNum
			case '-':
				result -= currentNum
			case '*':
				result *= currentNum
			case '/':
				if currentNum == 0 {
					return "", ErrInvalidExpression
				}
				result /= currentNum
			}
			currentNum = 0
			currentOp = ch
		}
	}

	return strconv.Itoa(result), nil
}
