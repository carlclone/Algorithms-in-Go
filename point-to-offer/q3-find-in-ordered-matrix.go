package point_to_offer

import "fmt"

/*
本题分析重点,从具体问题入手,然后用文字描述出解题过程的完整逻辑
 较复杂问题,用具体的例子找出规律的能力
*/
//TODO;文字描述逻辑

func main() {
	matrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	res := find(matrix, len(matrix), len(matrix[0]), 5)
	fmt.Println(res)

}

func find(matrix [][]int, rows, cols int, target int) bool {

	if rows == 0 || cols == 0 { //写的时候漏了
		return false
	}

	j := 0
	k := cols - 1

	for j <= rows-1 && k >= 0 {
		val := matrix[j][k]
		if val > target {
			k--
			continue
		}
		if val < target {
			j++
			continue
		}
		return true
	}
	return false
}
