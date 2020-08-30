package main

/*
todo;
改动迷宫, 改动起点终点测试
打印出路径

*/

import (
	"fmt"
	"os"
)

func main() {
	// imooc google engineer go course
	//https://pan.baidu.com/disk/home#/all?vmode=list&path=%2F%E7%AC%AC12%E7%AB%A0%20%E8%BF%B7%E5%AE%AB%E7%9A%84%E5%B9%BF%E5%BA%A6%E4%BC%98%E5%85%88%E6%90%9C%E7%B4%A2
	//先处理输入文件,得到 maze 数组[][]int
	maze := readMaze("application-algo/bfs-maze-walk/maze.in")

	start := point{0, 0}
	end := point{5, 4}
	res := walk(maze, start, end)
	for _, row := range res {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
	fmt.Println()
	//将maze 传入 solve 进行迷宫处理
	//采取 bfs 的策略,上下左右进行探索,case 分析
	//创建一个和 maze 相同的二维数组,存储探索的结果,是走还是不走 , 是否已经 visit 过了 , 是否死胡同,应该停下来了

	//思维学习
	// 走迷宫最短路径问题 , 转换成了另一个问题 :走到某个位置需要的最少步数(step) , 形成一个步数的二维数组,
	// 再从终点往回追溯得到一条路径

	//语言技能 fscanf 处理输入数据 ,slice 实现队列(也可以模拟栈) ,
	// 循环创建二位 slice(初始化) , slice 理解为slice结构体,所以第一维分配的是多个slice 结构体(len+底层数组指针)的空间 , 第二位是分配的每个 slice 底层实际的数组空间
	// 抽象出 point 的概念(ood)

	//画图
	//定义
	//伪代码
	//基本功能
	//边界条件,撞墙
	//异常处理:

}

type point struct {
	i, j int
}

//不变量 , cs61a 和 eloquent js 讲到 , 提高可维护性的手段
//到下一个新位置
func (p point) to(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p *point) equal(r point) bool {
	return p.i == r.i && p.j == r.j
}

//取二维数组 在该点的 val , 处理越界
func (p *point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i > len(grid)-1 {
		return 0, false
	}
	if p.j < 0 || p.j > len(grid[0])-1 {
		return 0, false
	}

	return grid[p.i][p.j], true
}

var directions = [4]point{
	{-1, 0}, //上,逆时针
	{0, -1}, //左
	{1, 0},
	{0, 1},
}

//0表示没走过 , >0 的表示走过或者是墙
func walk(maze [][]int, start point, end point) [][]int {
	//存储走到某个格式需要的最少步数 , bfs 和动态规划 , 最优解之间有什么关系
	// (都是穷举式的搜索? 但 dp 用了 memo 优化)
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//图的 bfs , 想想tree 的 bfs
	queue := []point{}
	queue = append(queue, start)

	for len(queue) > 0 {
		curPos := queue[0]
		queue = queue[1:]

		if curPos.equal(end) {
			break
		}

		for _, dir := range directions {
			nextPos := curPos.to(dir)
			// next 是 0 (maze 和 steps 里) 才可以走过去
			val1, ok1 := nextPos.at(maze)
			if !ok1 || val1 != 0 {
				continue
			}
			val2, ok2 := nextPos.at(steps)
			if !ok2 || val2 != 0 {
				continue
			}
			// 不是起点 !=start
			if nextPos.equal(start) {
				continue
			}

			//填格子
			curVal, _ := curPos.at(steps)
			steps[nextPos.i][nextPos.j] = curVal + 1

			//bfs , 把子节点放入队列
			queue = append(queue, nextPos)
		}
	}

	return steps
}

func readMaze(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	file.Close()

	return maze

}
