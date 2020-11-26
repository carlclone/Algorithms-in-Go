package array_and_hashtable

/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

func spiralOrder(matrix [][]int) []int {
	//coord := [][]int{ {0, 1}, {1, 0}, {0, -1},{-1, 0}} // right down  left up
	if len(matrix) == 0 {
		return []int{}
	}
	//start:=[]int{0,0}
	res := []int{}
	resCount := 0

	n := len(matrix)
	m := len(matrix[0])

	uw := 0
	rw := 0
	dw := 0
	lw := 0

	curPosH := 0
	curPosV := -1

	for {
		for i := 1; i <= m-rw-lw; i++ {
			curPosV++
			meta := matrix[curPosH][curPosV]
			res = append(res, meta)
			resCount++
		}
		uw += 1
		if resCount == m*n {
			break
		}

		for i := 1; i <= n-uw-dw; i++ {
			curPosH++
			meta := matrix[curPosH][curPosV]
			res = append(res, meta)
			resCount++
		}
		rw += 1
		if resCount == m*n {
			break
		}

		for i := 1; i <= m-rw-lw; i++ {
			curPosV--
			meta := matrix[curPosH][curPosV]
			res = append(res, meta)
			resCount++
		}
		dw += 1
		if resCount == m*n {
			break
		}

		for i := 1; i <= n-dw-uw; i++ {
			curPosH--
			meta := matrix[curPosH][curPosV]
			res = append(res, meta)
			resCount++
		}
		lw += 1
		if resCount == m*n {
			break
		}

	}

	return res
}

// version 2

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	startRow, endRow := 0, len(matrix)-1
	startCol, endCol := 0, len(matrix[0])-1

	dir := 0
	res := []int{}
	for startRow <= endRow && startCol <= endCol {
		switch dir {
		case 0: //RIGHT
			for col := startCol; col <= endCol; col++ {
				res = append(res, matrix[startRow][col])
			}
			startRow++
			break
		case 1: //DOWN
			for row := startRow; row <= endRow; row++ {
				res = append(res, matrix[row][endCol])
			}
			endCol--
			break
		case 2: //LEFT
			for col := endCol; col >= startCol; col-- {
				res = append(res, matrix[endRow][col])
			}
			endRow--
			break
		case 3: //UP
			for row := endRow; row >= startRow; row-- {
				res = append(res, matrix[row][startCol])
			}
			startCol++
			break

		}
		dir = (dir + 1) % 4
	}
	return res
}
