package array_and_hashtable

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	origin := []int{}
	for _, v := range nums {
		origin = append(origin, v)
	}
	sort.Ints(nums)
	res := []int{}
	for k, _ := range origin {
		if nums[k] != origin[k] {
			res = append(res, k)
		}
	}

	if len(res) > 0 {
		return res[len(res)-1] - res[0] + 1
	}
	return 0
}
