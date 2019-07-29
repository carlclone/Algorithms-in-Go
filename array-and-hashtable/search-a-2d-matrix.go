package array_and_hashtable

/*
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
Example 1:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
Example 2:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	currRow := 0
	currCol := len(matrix[0]) - 1

	for inArea(currRow, currCol, matrix) {
		if target > matrix[currRow][currCol] {
			currRow += 1
			continue
		}
		if target < matrix[currRow][currCol] {
			currCol -= 1
			continue
		}
		if target == matrix[currRow][currCol] {
			return true
		}
	}

	return false
}

func inArea(row int, col int, matrix [][]int) bool {
	return row >= 0 && row <= len(matrix)-1 && col >= 0 && col <= len(matrix[0])-1
}
