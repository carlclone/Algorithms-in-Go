package array_and_hashtable

//Given an unsorted integer array, find the smallest missing positive integer.
//
//Example 1:
//
//Input: [1,2,0]
//Output: 3
//Example 2:
//
//Input: [3,4,-1,1]
//Output: 2
//Example 3:
//
//Input: [7,8,9,11,12]
//Output: 1
//Note:
//
//Your algorithm should run in O(n) time and uses constant extra space.

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	map1 := make(map[int]bool)
	max := nums[0]
	for _, v := range nums {
		map1[v] = true
		if v > max {
			max = v
		}
	}
	if max < 0 {
		max = 0
	}

	for i := 1; i <= max; i++ {
		if _, ok := map1[i]; !ok {
			return i
		}
	}
	return max + 1

}
