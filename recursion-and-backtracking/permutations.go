package recursion_and_backtracking

/*
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

func permute(nums []int) [][]int {
	var res [][]int
	lenNums := len(nums)
	map1 := map[int]bool{}
	for k, v := range nums {
		tmp := []int{v}
		map1[k] = true
		p(k, v, tmp, map1, lenNums, nums, &res)
		map1[k] = false

	}
	return res

}

func p(k int, v int, t []int, map1 map[int]bool, lenNums int, nums []int, res *[][]int) {
	if len(t) == lenNums {
		*res = append(*res, t)
		return
	}
	for k, v := range nums {
		if map1[k] == true {
			continue
		}
		tmp := append(t, v)
		map1[k] = true
		p(k, v, tmp, map1, lenNums, nums, res)
		map1[k] = false
	}
}
