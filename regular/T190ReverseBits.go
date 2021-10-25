package regular

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	count := 0
	for count < 32 {
		remain := num & 0x1
		res = res<<1 + remain
		num >>= 1
		count++
	}
	return res
}
