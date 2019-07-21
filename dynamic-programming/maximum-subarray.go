package dynamic_programming

//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
//Example:
//
//Input: [-2,1,-3,4,-1,2,1,-5,4],
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
//Follow up:
//
//If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	maxCurr := nums[0]
	currSubArrVal := nums[0]
	for i := 1; i < len(nums); i++ {
		//fmt.Println(maxCurr)
		//fmt.Println(currSubArrVal)

		if currSubArrVal >= 0 && nums[i] >= 0 {
			maxCurr = max(maxCurr, currSubArrVal+nums[i])
			currSubArrVal = maxCurr
		} else {
			maxCurr = max(maxCurr, nums[i])
			currSubArrVal = maxCurr + nums[i]
		}
	}
	return maxCurr
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
