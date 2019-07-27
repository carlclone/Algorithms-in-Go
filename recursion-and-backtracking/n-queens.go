package recursion_and_backtracking

import "fmt"

/*
Example:

Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
*/

func solveNQueens(n int) [][]string {
	s := Solution{col: make(map[int]bool), dia1: make(map[int]bool), dia2: make(map[int]bool)}
	s.putQueen(n, 0)
	fmt.Println(len(s.res))
	out := [][]string{}
	for _, v := range s.res {
		out = append(out, gen(v, n))
	}
	return out
}

func totalNQueens(n int) int {
	s := Solution{col: make(map[int]bool), dia1: make(map[int]bool), dia2: make(map[int]bool)}
	s.putQueen(4, 0)
	return len(s.res)
}

type Solution struct {
	res    [][]int
	col    map[int]bool
	dia1   map[int]bool
	dia2   map[int]bool
	posArr []int
}

func (t *Solution) putQueen(n int, curRow int) {
	if curRow == n {
		c := []int{}
		for i := 0; i < n; i++ {
			c = append(c, i)
		}
		copy(c, t.posArr)
		t.res = append(t.res, c)
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
			t.posArr = t.posArr[:len(t.posArr)-1]
		}
	}
}

func gen(posArr []int, n int) []string {
	res := []string{}

	for _, v := range posArr {
		str := ""
		for i := 0; i <= n-1; i++ {
			if i == v {
				str += "Q"
			} else {
				str += "."
			}
		}
		res = append(res, str)
	}
	fmt.Println(res)
	return res
}
