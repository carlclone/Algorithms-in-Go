package stack

func isValid(s string) bool {
	/*
		iterate all if match pop , else insert
	*/
	stack := []rune{}
	maps := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack = append(stack, ' ')
	for _, v := range s {
		if maps[v] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 1
}
