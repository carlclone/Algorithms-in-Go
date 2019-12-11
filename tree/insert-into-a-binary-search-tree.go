package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	insertInto(root, val)
	return root
}
func insertInto(root *TreeNode, val int) {
	if val < root.Val && root.Left != nil {
		insertInto(root.Left, val)
	}
	if val > root.Val && root.Right != nil {
		insertInto(root.Right, val)
	}
	if val < root.Val && root.Left == nil {
		root.Left = &TreeNode{Val: val}
	}

	if val > root.Val && root.Right == nil {
		root.Right = &TreeNode{Val: val}
	}
}
