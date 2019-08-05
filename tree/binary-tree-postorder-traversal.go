package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Command struct {
	Type string // go , print
	Val  *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := NewStack()
	if root == nil {
		return res
	}
	/*

	 */
	stack.push(&Command{Type: "print", Val: root})
	if root.Right != nil {
		stack.push(&Command{Type: "go", Val: root.Right})
	}
	if root.Left != nil {
		stack.push(&Command{Type: "go", Val: root.Left})
	}

	for !stack.isEmpty() {
		curr := stack.pop()
		if curr.Type == "go" {
			stack.push(&Command{Type: "print", Val: curr.Val})
			if curr.Val.Right != nil {
				stack.push(&Command{Type: "go", Val: curr.Val.Right})
			}
			if curr.Val.Left != nil {
				stack.push(&Command{Type: "go", Val: curr.Val.Left})
			}
		} else if curr.Type == "print" {
			res = append(res, curr.Val.Val)
		}
	}
	return res

}

type Stack struct {
	data [1000]*Command
	pos  int
}

func (s *Stack) pop() *Command {
	s.pos--
	return s.data[s.pos]

}

func (s *Stack) push(t *Command) {
	s.data[s.pos] = t
	s.pos++
}

func (s *Stack) isEmpty() bool {
	return s.pos == 0
}

func NewStack() *Stack {
	return &Stack{}
}
