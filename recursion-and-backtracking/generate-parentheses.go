package recursion_and_backtracking

func generateParenthesis(n int) []string {
	list := []string{}
	gen(0, 0, n, &list, "")
	return list
}

func gen(left int, right int, n int, list *[]string, result string) {
	//这个判断把全部不符合的都排除掉了 , 简直了
	if left < right {
		return
	}

	if left == n && right == n {
		//append to
		*list = append(*list, result)
	}
	if left < n {
		gen(left+1, right, n, list, result+"(")
	}
	if right < n {
		gen(left, right+1, n, list, result+")")
	}
}
