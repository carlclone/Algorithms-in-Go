package point_to_offer

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	max := dp[0]
	//        nums[i]      , dp[0]+nums[1] < nums[1]
	//f(i)
	//        dp[i-1]+nums[i] , dp[0]+nums[1] >= nums[1]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] < nums[i] {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
