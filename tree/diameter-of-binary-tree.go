package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
情况讨论:
最长的长度 = max(每个节点的最长长度)
每个节点的最长长度= 左节点的max + 右节点的max +1
某个节点的最长长度=max(左节点最长,右节点最长)+1

从关联问题max depth of a node出发 , 遍历所有节点 ,  得出答案
还有重叠子问题 , 减少了重复求解 , 这解法浓缩了精华

*/
type Solution struct {
	ans int
}

func (t *Solution) depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := t.depth(node.Left)
	rightDepth := t.depth(node.Right)
	if leftDepth+rightDepth+1 > t.ans {
		t.ans = leftDepth + rightDepth + 1
	}
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	s := &Solution{ans: 1}
	s.depth(root)
	return s.ans - 1
}
