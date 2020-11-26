package dynamic_programming

//Say you have an array-and-hashtable for which the ith element is the price of a given stock on day i.
//
//If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
//
//Note that you cannot sell a stock before you buy one.
//
//Example 1:
//
//Input: [7,1,5,3,6,4]
//Output: 5
//Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//             Not 7-1 = 6, as selling price needs to be larger than buying price.
//Example 2:
//
//Input: [7,6,4,3,1]
//Output: 0
//Explanation: In this case, no transaction is done, i.e. max profit = 0.

func maxProfit(prices []int) int {
	var (
		leftMinIndexArr []int
		price           int
		index           int
		maxProfit       int
	)
	leftMinIndexArr = make([]int, len(prices))
	for index, price = range prices {
		if index == 0 {
			leftMinIndexArr[index] = price
			continue
		}
		if price > leftMinIndexArr[index-1] {
			leftMinIndexArr[index] = leftMinIndexArr[index-1]
		} else {
			leftMinIndexArr[index] = price
		}
	}

	for index, price = range prices {
		if price-leftMinIndexArr[index] > maxProfit {
			maxProfit = price - leftMinIndexArr[index]
		}
	}
	return maxProfit
}

// 股票系列通用 DP 解法
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := make([][][]int, len(prices))
	for k, _ := range profit {
		profit[k] = make([][]int, 2)
		for j, _ := range profit[k] {
			profit[k][j] = make([]int, 2)
		}
	}
	profit[0][0][0], profit[0][0][1] = 0, -prices[0]
	profit[0][1][0], profit[0][1][1] = 0, 0

	i := 1
	for i < len(prices) {
		profit[i][0][0] = profit[i-1][0][0]
		profit[i][0][1] = max(profit[i-1][0][1], profit[i-1][0][0]-prices[i])

		profit[i][1][0] = max(profit[i-1][1][0], profit[i-1][0][1]+prices[i])
		i++
	}
	end := len(prices) - 1
	return max(profit[end][0][0], profit[end][1][0])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// clear solution
func maxProfit(prices []int) int {
	// write code here
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxPro := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > maxPro {
			maxPro = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxPro
}
