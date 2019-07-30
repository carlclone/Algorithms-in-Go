package tree

/*
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

*/

func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumCustom(root, sum)
}

func hasPathSumCustom(node *TreeNode, sum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil && node.Val-sum == 0 {
		return true
	} else if node.Left == nil && node.Right == nil && node.Val-sum != 0 {
		return false
	} else {
		res1 := false
		res2 := false

		if node.Left != nil {
			res1 = hasPathSumCustom(node.Left, sum-node.Val)
		}
		if node.Right != nil {
			res2 = hasPathSumCustom(node.Right, sum-node.Val)
		}

		if res1 || res2 {
			return true
		} else {
			return false
		}
	}

}
