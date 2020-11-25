package main

/**
 *
 * @param A int整型一维数组
 * @param target int整型
 * @return int整型
 */
func search(A []int, target int) int {
	// write code here
	l, r := 0, len(A)

	for l <= r {
		mid := (l + r) / 2

		if A[mid] == target {
			return mid
		}

		if A[mid] > A[len(A)-1] {
			if target > A[0] && target < A[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target < A[len(A)-1] && target > A[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
