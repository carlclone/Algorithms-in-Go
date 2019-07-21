package dynamic_programming

import "fmt"

//You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
//
//Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
//
//Example 1:
//
//Input: [1,2,3,1]
//Output: 4
//Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//             Total amount you can rob = 1 + 3 = 4.
//Example 2:
//
//Input: [2,7,9,3,1]
//Output: 12
//Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
//             Total amount you can rob = 2 + 9 + 1 = 12.

type Solution struct {
	memo map[int]int
	len  int
	nums []int
}

func (t *Solution) rob(n int) int {
	// n>len return 0
	if n > t.len-1 {
		return 0
	}

	max := -1

	for i := n; i <= t.len-1; i++ {
		if _, ok := t.memo[i+2]; ok {
			max = maxC(max, t.nums[i]+t.memo[i+2])
			continue
		}

		cur := t.nums[i] + t.rob(i+2)

		max = maxC(max, cur)
	}

	t.memo[n] = max
	return max
}

func rob(nums []int) int {
	s := Solution{memo: make(map[int]int), len: len(nums), nums: nums}
	res := s.rob(0)
	return res
}

func maxC(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func p(v interface{}) {
	fmt.Println(v)
}
