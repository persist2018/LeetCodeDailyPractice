package main

// stack pop
func pop(stack []rune) (stackAfterPop []rune, char rune) {
	// fmt.Println("stack", stack)
	if len(stack) == 0 {
		return []rune{}, ' '
	}
	if len(stack)-1 == 0 {
		return []rune{}, stack[0]
	}
	return stack[0 : len(stack)-1], stack[len(stack)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minThree(a, b, c int) int {
	temp := 0
	if a > b {
		temp = b
	} else {
		temp = a
	}
	if c > temp {
		return temp
	}
	return c
}
