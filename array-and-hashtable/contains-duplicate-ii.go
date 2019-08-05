package array_and_hashtable

//Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.
//
//Example 1:
//
//Input: nums = [1,2,3,1], k = 3
//Output: true
//Example 2:
//
//Input: nums = [1,0,1,1], k = 1
//Output: true
//Example 3:
//
//Input: nums = [1,2,3,1,2,3], k = 2
//Output: false

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	if k == 0 {
		return false
	}
	map1 := make(map[int]int)
	l := 0
	r := 0
	map1[nums[0]] = 0

	for r < len(nums)-1 {
		inMap := false
		if _, ok := map1[nums[r+1]]; ok {
			inMap = true
		}

		if r+1-l <= k && inMap {
			return true
		} else {
			if r-l < k-1 {
				r++
				map1[nums[r]] = r
				continue
			}
			if r-l == k-1 {
				delete(map1, nums[l])
				l++

				r++
				map1[nums[r]] = r
				continue
			}
		}
	}

	return false

}
