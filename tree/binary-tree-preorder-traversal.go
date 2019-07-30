package tree

/*
Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

//迭代解法
func preorderTraversal2(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}
	stack := StackNode{}
	stack.Push(root)
	for !stack.IsEmpty() {
		cur := stack.Pop()
		r = append(r, cur.Val)
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}
	return r
}

type StackNode struct {
	data [1000]*TreeNode
	pos  int
}

func NewStackNode() *StackNode {
	return &StackNode{}
}
func (s *StackNode) IsEmpty() bool {
	return s.pos == 0
}
func (s *StackNode) Push(n *TreeNode) {
	s.data[s.pos] = n
	s.pos++
}
func (s *StackNode) Pop() (n *TreeNode) {
	s.pos--
	return s.data[s.pos]
}

//递归解法
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal1(root *TreeNode) []int {
	res := []int{}
	pre(root, &res)

	return res
}

func pre(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	//fmt.Println(res)
	pre(node.Left, res)
	pre(node.Right, res)
}
