package array_and_hashtable

//Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
//
//You may return any answer array that satisfies this condition.
//
//
//
//Example 1:
//
//Input: [3,1,2,4]
//Output: [2,4,3,1]
//The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
//
//
//Note:
//
//1 <= A.length <= 5000
//0 <= A[i] <= 5000

//额外空间
func sortArrayByParity1(A []int) []int {
	len := len(A)
	if len == 0 || len == 1 {
		return A
	}

	new := make([]int, len)
	p1 := 0
	p2 := len - 1
	for i := 0; i <= len-1; i++ {
		if A[i]%2 == 0 {
			new[p1] = A[i]
			p1++
		} else {
			new[p2] = A[i]
			p2--
		}
	}
	return new

}

//不使用额外空间
func sortArrayByParity(A []int) []int {
	len := len(A)
	if len == 0 || len == 1 {
		return A
	}

	//p1 := 0
	p2 := -1

	for p2+1 <= len-1 {
		if A[p2+1]%2 == 0 {
			p2++
		} else {
			p3 := p2 + 1
			for i := p3; i <= len-1; i++ {
				if A[i]%2 == 0 {
					tmp := A[p3]
					A[p3] = A[i]
					A[i] = tmp
					p2++
					break
				}
				//优化 , 当p3往后遍历都没有奇数的时候 , 直接返回
				if i == len-1 && A[i]%2 == 1 {
					return A
				}
			}
		}
	}
	return A
}
