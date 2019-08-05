package array_and_hashtable

//Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
//
//Integers in each row are sorted in ascending from left to right.
//Integers in each column are sorted in ascending from top to bottom.
//Example:
//
//Consider the following matrix:
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//Given target = 5, return true.
//
//Given target = 20, return false.

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
