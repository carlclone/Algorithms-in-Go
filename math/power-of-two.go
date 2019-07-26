package math

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}
	for {
		if n%2 == 0 {
			n = n / 2
			if n == 1 {
				return true
			}
		} else {
			return false
		}
	}
}
