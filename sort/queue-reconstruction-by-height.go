package sort

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// sort input based on h first then k
	// if h is equal, place element with bigger k value first
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	head := &Node{Val: people[0]}

	for k, v := range people {
		if k == 0 {
			continue
		}
		putIntoList(&head, v)
	}

	res := iterateListGetArr(&head)
	return res
}

type Node struct {
	Val  []int
	Next *Node
}

func main() {
	arr := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	fmt.Println(reconstructQueue(arr))
}
