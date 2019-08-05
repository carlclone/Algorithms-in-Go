package array_and_hashtable

//Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.
//
//Example 1:
//
//Input: nums = [1,2,3,1], k = 3, t = 0
//Output: true
//Example 2:
//
//Input: nums = [1,0,1,1], k = 1, t = 2
//Output: true
//Example 3:
//
//Input: nums = [1,5,9,1,5,9], k = 2, t = 3
//Output: false

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
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
		exist := false
		for k, _ := range map1 {
			if math.Abs(float64(nums[r+1]-k)) <= float64(t) {
				exist = true
			}
		}

		if r+1-l <= k && exist {
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
