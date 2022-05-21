package main

// import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)
	c := '('
	for _, bracket := range s {
		switch bracket {
		case '(':
			stack = append(stack, bracket)
		case '[':
			stack = append(stack, bracket)
		case '{':
			stack = append(stack, bracket)
		case ')':
			stack, c = pop(stack)
			if c != '(' {
				return false
			}
		case ']':
			stack, c = pop(stack)
			if c != '[' {
				return false
			}
		case '}':
			stack, c = pop(stack)
			if c != '{' {
				return false
			}

		}
	}
	return len(stack) == 0
}
