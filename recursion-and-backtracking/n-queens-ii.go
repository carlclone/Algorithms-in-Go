package recursion_and_backtracking

/*
Example:

Input: 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
*/

func totalNQueens(n int) int {
	s := Solution{col: make(map[int]bool), dia1: make(map[int]bool), dia2: make(map[int]bool)}
	s.putQueen(n, 0)
	return len(s.res)
}

//func main() {
//	s:=Solution{col:make(map[int]bool),dia1:make(map[int]bool),dia2:make(map[int]bool)}
//	s.putQueen(4,0)
//	return len(s.res)
//}

type Solution struct {
	res    [][]int
	col    map[int]bool
	dia1   map[int]bool
	dia2   map[int]bool
	posArr []int
}

func (t *Solution) putQueen(n int, curRow int) {
	if curRow == n {
		t.res = append(t.res, t.posArr)

		t.posArr = []int{}
		return
	}

	for i := 0; i <= n-1; i++ {
		d1 := i + curRow
		d2 := i - curRow + n - 1
		if t.col[i] != true && t.dia1[d1] != true && t.dia2[d2] != true {
			t.posArr = append(t.posArr, i)
			t.col[i] = true
			t.dia1[d1] = true
			t.dia2[d2] = true
			t.putQueen(n, curRow+1)
			t.col[i] = false
			t.dia1[d1] = false
			t.dia2[d2] = false
		}
	}
	t.posArr = []int{}
}
