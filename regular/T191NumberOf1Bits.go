package regular

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num&0x1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}
