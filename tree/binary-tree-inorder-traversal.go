package tree

/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// using recursion
func inorderTraversal(root *TreeNode) []int {
	var (
		res []int
	)

	inorder(&res, root)
	return res
}

func inorder(res *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	inorder(res, node.Left)
	*res = append(*res, node.Val)
	inorder(res, node.Right)
}

type Command struct {
	Type string
	Val  *TreeNode
}

//using stack emulation
func inorderTraversalStack(root *TreeNode) (res []int) {
	var (
		stack   []*Command // print , go
		pop     *TreeNode
		command *Command
	)

	if root == nil {
		return
	}
	command = &Command{Type: "go", Val: root}
	stack = append(stack, command)

	for len(stack) != 0 {
		command = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if command.Type == "go" {
			pop = command.Val
			if pop.Right != nil {
				stack = append(stack, &Command{Type: "go", Val: pop.Right})
			}

			stack = append(stack, &Command{Type: "print", Val: pop})
			if pop.Left != nil {
				stack = append(stack, &Command{Type: "go", Val: pop.Left})
			}
		}
		if command.Type == "print" {
			res = append(res, command.Val.Val)
		}
	}
	return
}
