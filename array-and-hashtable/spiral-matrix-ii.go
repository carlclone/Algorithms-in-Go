package array_and_hashtable

/*
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	if n == 0 {
		return res
	}
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	resCount := 0

	m := n

	uw := 0
	rw := 0
	dw := 0
	lw := 0

	curPosH := 0
	curPosV := -1

	curVal := 1

	for {
		for i := 1; i <= m-rw-lw; i++ {
			curPosV++
			res[curPosH][curPosV] = curVal
			curVal++
			resCount++
		}
		uw += 1
		if resCount == m*n {
			break
		}

		for i := 1; i <= n-uw-dw; i++ {
			curPosH++
			res[curPosH][curPosV] = curVal
			curVal++
			resCount++
		}
		rw += 1
		if resCount == m*n {
			break
		}

		for i := 1; i <= m-rw-lw; i++ {
			curPosV--
			res[curPosH][curPosV] = curVal
			curVal++
			resCount++
		}
		dw += 1
		if resCount == m*n {
			break
		}

		for i := 1; i <= n-dw-uw; i++ {
			curPosH--
			res[curPosH][curPosV] = curVal
			curVal++
			resCount++
		}
		lw += 1
		if resCount == m*n {
			break
		}

	}

	return res
}
