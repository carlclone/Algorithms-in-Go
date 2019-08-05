package tree

//Invert a binary tree.
//
//Example:
//
//Input:
//
//     4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//Output:
//
//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//Trivia:
//This problem was inspired by this original tweet by Max Howell:
//
//Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.

//队列,BFS
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			if node.Left.Left != nil || node.Left.Right != nil {
				queue = append(queue, node.Left)
			}
		}
		if node.Right != nil {
			if node.Right.Left != nil || node.Right.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
	}
	return root
}

//递归
