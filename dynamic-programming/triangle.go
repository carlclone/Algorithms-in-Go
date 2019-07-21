package dynamic_programming

//Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
//
//For example, given the following triangle
//
//[
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
//The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
//
//Note:
//
//Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

type Solution struct {
	memo        [][]int
	bottomIndex int
	arr         [][]int
}

func (t *Solution) findMin(ver int, hor int) int {
	if ver == t.bottomIndex {
		return t.arr[ver][hor]
	}
	if t.memo[ver][hor] != 0 {
		return t.memo[ver][hor]
	}

	res1 := t.findMin(ver+1, hor)
	res2 := t.findMin(ver+1, hor+1)

	res := t.arr[ver][hor]
	if res1 > res2 {
		res += res2
	} else {
		res += res1
	}
	t.memo[ver][hor] = res
	return res
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	memo := make([][]int, len(triangle))
	for k, _ := range memo {
		memo[k] = make([]int, len(triangle))

	}
	s := Solution{memo: memo, bottomIndex: len(triangle) - 1, arr: triangle}
	return s.findMin(0, 0)
}
