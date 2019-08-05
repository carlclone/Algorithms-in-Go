package dynamic_programming

//Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
//
//Example 1:
//
//Input:
//11110
//11010
//11000
//00000
//
//Output: 1
//Example 2:
//
//Input:
//11000
//11000
//00100
//00011
//
//Output: 3

type Solution struct {
	visited map[int]map[int]bool
}

func (t *Solution) findCloseGrids(grid [][]byte, srow int, scol int) int {
	if grid[srow][scol] != '1' {
		return 0
	}

	res := 1
	if _, ok := t.visited[srow]; !ok {
		t.visited[srow] = make(map[int]bool)
	}
	t.visited[srow][scol] = true

	coord := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //up right down  left

	for _, v := range coord {
		nrow := srow + v[0]
		ncol := scol + v[1]
		if inArea(ncol, nrow, grid) && !t._visited(ncol, nrow) {
			res += t.findCloseGrids(grid, nrow, ncol)
		}
	}
	return res
}

func (t *Solution) _visited(col int, row int) bool {
	if _, ok := t.visited[row][col]; ok {
		return true

	} else {
		return false
	}
}

func inArea(col int, row int, board [][]byte) bool {
	return row >= 0 && row <= len(board)-1 && col >= 0 && col <= len(board[0])-1
}

func numIslands(grid [][]byte) int {
	s := Solution{visited: make(map[int]map[int]bool)}
	res := 0
	for row, rowdata := range grid {
		for col, _ := range rowdata {
			if !s._visited(col, row) {
				if s.findCloseGrids(grid, row, col) > 0 {
					res += 1
				}
			}
		}
	}
	return res
}
