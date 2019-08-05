package tree

import "strconv"

//Given a binary tree, return all root-to-leaf paths.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Input:
//
//   1
// /   \
//2     3
// \
//  5
//
//Output: ["1->2->5", "1->3"]
//
//Explanation: All root-to-leaf paths are: 1->2->5, 1->3

//TODO;代码可以精简优化
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	s := Solution{}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	if root.Left != nil {
		s.treePathsIn(root.Left, strconv.Itoa(root.Val))
	}
	if root.Right != nil {
		s.treePathsIn(root.Right, strconv.Itoa(root.Val))
	}
	return s.res
}

type Solution struct {
	res []string
}

func (t *Solution) treePathsIn(node *TreeNode, rootPath string) {
	if node.Left == nil && node.Right == nil {
		t.res = append(t.res, rootPath+"->"+strconv.Itoa(node.Val))
		return
	} else {
		if node.Left != nil {
			t.treePathsIn(node.Left, rootPath+"->"+strconv.Itoa(node.Val))
		}
		if node.Right != nil {
			t.treePathsIn(node.Right, rootPath+"->"+strconv.Itoa(node.Val))
		}
	}
}
