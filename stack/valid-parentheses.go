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

//version 2

type Stack struct {
	list []rune
}

func (t *Stack) Size() int {
	return len(t.list)
}

func (t *Stack) Pop() rune {
	res := t.list[len(t.list)-1]
	t.list = t.list[:len(t.list)-1]
	return res
}

func (t *Stack) Peek() rune {
	res := t.list[len(t.list)-1]

	return res
}

func (t *Stack) Push(val rune) {
	t.list = append(t.list, val)

}

/**
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isValid(s string) bool {
	// write code here
	if len(s) == 0 {
		return false
	}
	stack := &Stack{}
	r := []rune(s)
	stack.Push(r[0])
	for i := 1; i < len(s); i++ {
		if (stack.Peek() == '(' && r[i] == ')') ||
			(stack.Peek() == '{' && r[i] == '}') ||
			(stack.Peek() == '[' && r[i] == ']') {
			stack.Pop()
		} else {
			stack.Push(r[i])
		}

	}

	if stack.Size() == 0 {
		return true
	}
	return false

}
