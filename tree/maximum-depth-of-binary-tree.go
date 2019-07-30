package tree

/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftD := maxDepth(root.Left)
	rightD := maxDepth(root.Right)

	if leftD > rightD {
		return leftD + 1
	}
	return rightD + 1
}
