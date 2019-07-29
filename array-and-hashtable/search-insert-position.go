package array_and_hashtable

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0
*/

func searchInsert(nums []int, target int) int {
	l, len, r := 0, len(nums), len(nums)-1

	for l <= r {
		mid := (r + l) / 2
		if target == nums[mid] {
			return mid
		}
		if mid == 0 && target < nums[mid] {
			return 0
		}
		if mid == len-1 && target > nums[mid] {
			return len
		}
		if target < nums[mid] && target > nums[mid-1] {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
			continue
		}
		if target < nums[mid] {
			r = mid - 1
			continue
		}
	}
	return 0
}
