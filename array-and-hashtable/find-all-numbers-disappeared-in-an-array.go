package array_and_hashtable

//Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
//
//Find all the elements of [1, n] inclusive that do not appear in this array.
//
//Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
//
//Example:
//
//Input:
//[4,3,2,7,8,2,3,1]
//
//Output:
//[5,6]

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	if len1 > len2 {
		match := true
		for k, _ := range str2 {
			if str1[k] != str2[k] {
				match = false
				break
			}
		}
		if match {
			str1Subs := str1[len2:]
			return gcdOfStrings(str1Subs, str2)
		}
	}

	if len1 < len2 {
		return gcdOfStrings(str2, str1)
	}

	if len1 == len2 {
		match := true
		for k, _ := range str2 {
			if str1[k] != str2[k] {
				match = false
				break
			}
		}
		if match {
			return str1
		}

	}

	return ""
}
