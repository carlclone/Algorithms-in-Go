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
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	in(&res, root)

	return res
}

func in(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	in(res, root.Left)
	*res = append(*res, root.Val)
	in(res, root.Right)
}
