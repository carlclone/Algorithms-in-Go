package dynamic_programming

func coinChange(coins []int, amount int) int {
	dp := make(map[int]int) // 金额为amount的时候使用的最少硬币数

	//初始化 , -1表示没有解

	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}
	dp[0] = 0 //

	//状态转移
	//dp[n] = min(for dp[n-coin]) +1
	maxint := 2147483647
	for i := 1; i <= amount; i++ {
		min := maxint
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i-coin]+1 < min && dp[i-coin] != -1 {
				min = dp[i-coin] + 1
				dp[i] = min
			}

		}
	}
	return dp[amount]
}
