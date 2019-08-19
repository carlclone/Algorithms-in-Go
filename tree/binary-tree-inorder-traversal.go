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

//using stack emulation
