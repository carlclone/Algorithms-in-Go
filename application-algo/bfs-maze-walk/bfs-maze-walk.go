package bfs_maze_walk

func main() {
	// imooc google engineer go course
	//https://pan.baidu.com/disk/home#/all?vmode=list&path=%2F%E7%AC%AC12%E7%AB%A0%20%E8%BF%B7%E5%AE%AB%E7%9A%84%E5%B9%BF%E5%BA%A6%E4%BC%98%E5%85%88%E6%90%9C%E7%B4%A2
	//先处理输入文件,得到 maze 数组[][]int
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
