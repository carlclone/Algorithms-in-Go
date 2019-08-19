package array_and_hashtable

//Given an array of integers, find if the array contains any duplicates.
//
//Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.
//
//Example 1:
//
//Input: [1,2,3,1]
//Output: true
//Example 2:
//
//Input: [1,2,3,4]
//Output: false
//Example 3:
//
//Input: [1,1,1,3,3,4,3,2,4,2]
//Output: true

// brute force
func containsDuplicate1(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}

	}
	return false
}

//using map
func containsDuplicate(nums []int) bool {
	var (
		freqMap map[int]int
		val     int
	)
	freqMap = make(map[int]int)
	for _, val = range nums {
		freqMap[val]++
		if freqMap[val] > 1 {
			return true
		}
	}

	return false
}
