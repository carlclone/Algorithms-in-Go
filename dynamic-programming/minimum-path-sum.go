package dynamic_programming

import "math"

//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
//
//Note: You can only move either down or right at any point in time.
//
//Example:
//
//Input:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//Output: 7
//Explanation: Because the path 1→3→1→1→1 minimizes the sum.

type Solution struct {
	memo          [][]int
	arr           [][]int
	verRightIndex int
	horRightIndex int
}

func (t *Solution) findMin(ver int, hor int) int {
	if ver == t.verRightIndex && hor == t.horRightIndex {
		return t.arr[ver][hor]
	}
	if ver > t.verRightIndex || hor > t.horRightIndex {
		return math.MaxInt32
	}
	if t.memo[ver][hor] != 0 {
		return t.memo[ver][hor]
	}

	res1 := t.findMin(ver+1, hor)
	res2 := t.findMin(ver, hor+1)

	res := t.arr[ver][hor]
	if res1 > res2 {
		res += res2
	} else {
		res += res1
	}
	t.memo[ver][hor] = res
	return res
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	memo := make([][]int, len(grid))
	for k, _ := range memo {
		memo[k] = make([]int, len(grid[0]))

	}
	s := Solution{memo: memo, verRightIndex: len(grid) - 1,
		horRightIndex: len(grid[0]) - 1, arr: grid}
	return s.findMin(0, 0)
}
