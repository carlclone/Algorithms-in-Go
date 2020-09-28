package byte_dance

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	str1 := strs[0]
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != str1[i] {
				return str1[:i]
			}
		}
	}
	return str1
}
