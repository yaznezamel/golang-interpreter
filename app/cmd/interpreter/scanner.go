package interpreter

import "fmt"

const EOF = "EOF"

// takes character as input and returns type as output
func TokenType(ch byte) (string, error) {

	switch ch {
	case '(':
		return "LEFT_PAREN", nil
	case ')':
		return "RIGHT_PAREN", nil

	case '{':
		return "LEFT_BRACE", nil
	case '}':
		return "RIGHT_BRACE", nil
	case '*':
		return "START", nil
	case '.':
		return "DOT", nil
	case ',':
		return "COMMA", nil
	case '+':
		return "PLUS", nil

	case '-':
		return "MINUS", nil
	default:
		return "", fmt.Errorf("unexpected character: %c", ch)
	}

}
