package array_and_hashtable

//Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).
//
//Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).
//
//Example:
//
//Input:
//[[0,0],[1,0],[2,0]]
//
//Output:
//2
//
//Explanation:
//The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i, v1 := range points {
		map1 := make(map[int]int)
		for j, v2 := range points {
			if i == j {
				continue
			}
			map1[dis(v1, v2)]++
		}

		for _, v := range map1 {
			res = res + v*(v-1)
		}
	}

	return res
}

func dis(point1 []int, point2 []int) int {
	return (point2[0]-point1[0])*(point2[0]-point1[0]) + (point2[1]-point1[1])*(point2[1]-point1[1])
}
