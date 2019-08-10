package array_and_hashtable

//https://leetcode.com/problems/pascals-triangle/

func generate(numRows int) [][]int {
	var (
		resArr     [][]int
		curRow     int
		firstIndex int
		lastIndex  int
		middles    int
	)

	resArr = make([][]int, numRows)
	for curRow = 0; curRow < numRows; curRow++ {
		firstIndex = 0
		lastIndex = curRow
		resArr[curRow] = make([]int, curRow+1)
		resArr[curRow][firstIndex] = 1
		resArr[curRow][lastIndex] = 1
		for middles = firstIndex + 1; middles < lastIndex; middles++ {
			resArr[curRow][middles] = resArr[curRow-1][middles-1] + resArr[curRow-1][middles]
		}
	}
	return resArr
}
