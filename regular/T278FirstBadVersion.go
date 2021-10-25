package regular

import "fmt"

func firstBadVersion(n int) int {
	hi := n
	lo := 1
	mid := 0
	for lo < hi {
		mid = (lo + hi) / 2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func isBadVersion(version int) bool {
	if version >= 2 {
		return true
	}
	return false
}

func T278() {
	res := firstBadVersion(3)
	fmt.Println(res)
}
