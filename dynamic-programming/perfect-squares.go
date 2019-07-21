package dynamic_programming

//Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.
//
//Example 1:
//
//Input: n = 12
//Output: 3
//Explanation: 12 = 4 + 4 + 4.
//Example 2:
//
//Input: n = 13
//Output: 2
//Explanation: 13 = 4 + 9.

func numSquares(n int) int {
	q := [][]int{}
	vm := make(map[int]int)
	q = append(q, []int{n, 0})

	for len(q) != 0 {
		//pop pairs
		pairs := q[0]
		q = q[1:]

		currNumRemain := pairs[0]
		steps := pairs[1]

		for i := 0; i*i <= currNumRemain; i++ {
			if currNumRemain-i*i == 0 {
				return steps + 1
			}
			if _, ok := vm[i*i]; ok {
				continue
			} else {
				q = append(q, []int{currNumRemain - i*i, steps + 1})
			}
		}
	}
	return 0
}
