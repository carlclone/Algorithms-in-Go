package recursion_and_backtracking

/*
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
*/

func exist(board [][]byte, word string) bool {
	for row, rowdata := range board {
		s := Solution{visited: make(map[int]map[int]bool)}
		res := false
		for col, _ := range rowdata {
			res = s.wordSearch(word, 0, len(word)-1, board, col, row)
			if res == true {
				return true
			}
		}
	}
	return false
}

type Solution struct {
	visited map[int]map[int]bool
}

func (t *Solution) wordSearch(word string, wordhead int, wordend int, board [][]byte, startcol int, startrow int) bool {
	if wordhead == len(word)-1 {
		return word[wordhead] == board[startrow][startcol]
	}

	if board[startrow][startcol] == word[wordhead] {
		coord := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //up right down  left
		if _, ok := t.visited[startrow]; !ok {
			t.visited[startrow] = make(map[int]bool)
		}
		t.visited[startrow][startcol] = true
		for _, v := range coord {
			newrow := startrow + v[0]
			newcol := startcol + v[1]
			if inArea(newcol, newrow, board) && !t._visited(newcol, newrow) {

				res := t.wordSearch(word, wordhead+1, wordend, board, newcol, newrow)
				if res == true {
					return true
				}

			}
		}
		delete(t.visited[startrow], startcol)
	}
	return false

}

func inArea(col int, row int, board [][]byte) bool {
	return row >= 0 && row <= len(board)-1 && col >= 0 && col <= len(board[0])-1
}

func (t *Solution) _visited(col int, row int) bool {
	if _, ok := t.visited[row][col]; ok {
		return true

	} else {
		return false
	}
}
