package tree

//Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.
//
//Example 1:
//
//Input:
//    5
//   / \
//  3   6
// / \   \
//2   4   7
//
//Target = 9
//
//Output: True
//
//
//Example 2:
//
//Input:
//    5
//   / \
//  3   6
// / \   \
//2   4   7
//
//Target = 28
//
//Output: False

func (t *Solution) findTarget1(root *TreeNode, k int) bool {
	node := root

	if node == nil {
		return false
	}
	if t.map1[k-node.Val] == true {
		return true
	} else {
		t.map1[node.Val] = true
	}
	res1 := t.findTarget1(root.Left, k)
	res2 := t.findTarget1(root.Right, k)
	if res1 || res2 {
		return true
	} else {
		return false
	}

}

type Solution struct {
	map1 map[int]bool
}

func findTarget(root *TreeNode, k int) bool {
	s := Solution{make(map[int]bool)}
	return s.findTarget1(root, k)
}
