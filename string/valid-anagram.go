package string

func isAnagram(s string, t string) bool {
	var (
		freqMap map[rune]int
		c       rune
		v       int
	)

	freqMap = make(map[rune]int)
	for _, c = range s {
		freqMap[c]++
	}

	for _, c = range t {
		if freqMap[c] <= 0 {
			return false
		} else {
			freqMap[c]--
		}
	}

	for _, v = range freqMap {
		if v > 0 {
			return false
		}
	}
	return true
}
