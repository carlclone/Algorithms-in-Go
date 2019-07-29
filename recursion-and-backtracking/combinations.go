package recursion_and_backtracking

import "fmt"

/*
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {
	res := [][]int{}

	for i := 1; i <= n; i++ {
		currA := []int{i}
		c(n, k, i, currA, &res)
	}
	return res

}

func c(n int, k int, currV int, currA []int, res *[][]int) {
	if len(currA) == k {
		tmp := make([]int, k)
		copy(tmp, currA)
		*res = append(*res, tmp)
		fmt.Println(tmp)
		//fmt.Println(res)
		fmt.Println(currA)
		return
	}
	for i := currV + 1; i <= n; i++ {
		currA := append(currA, i)
		c(n, k, i, currA, res)
	}
}
