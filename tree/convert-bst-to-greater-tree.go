package tree

func convertBST(root *TreeNode) *TreeNode {
	convert(root)
	return root
}

func convert(node *TreeNode) {
	if node == nil {
		return
	}

	leftOrigin := 0
	rightOrigin := 0
	nodeOrigin := node.Val

	if node.Left != nil {
		leftOrigin = node.Left.Val
	}
	if node.Right != nil {
		rightOrigin = node.Right.Val
	}

	if node.Left != nil {
		convert(node.Left)
		leftOrigin = node.Left.Val
		compareAndAdd(node.Left, rightOrigin, leftOrigin)
		compareAndAdd(node.Left, nodeOrigin, leftOrigin)
	}
	if node.Right != nil {
		convert(node.Right)
		rightOrigin = node.Right.Val
		compareAndAdd(node.Right, leftOrigin, rightOrigin)
		compareAndAdd(node.Right, nodeOrigin, rightOrigin)
	}
	compareAndAdd(node, leftOrigin, nodeOrigin)
	compareAndAdd(node, rightOrigin, nodeOrigin)
}

func compareAndAdd(node1 *TreeNode, v int, nodeOrigin int) {
	if v > nodeOrigin {
		node1.Val += v
	}
}
