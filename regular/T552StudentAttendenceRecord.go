package regular

type History struct {
	aCount       int
	endingLCount int
}

var cache = make(map[int]map[int]*History)

const MOD = 1000000007

var STATUS = []int{'A', 'L', 'P'}

func checkRecord(n int) int {
	if n == 0 {
		return 0
	}
	cache[0] = make(map[int]*History)
	cache[0]['A'] = &History{1, 0}
	cache[0]['L'] = &History{0, 1}
	cache[0]['P'] = &History{0, 0}
	count := 3
	for i := 1; i <= n; i++ {
		cache[i] = make(map[int]*History)
		for j := 0; j < len(STATUS); j++ {
			prev, prs := cache[i-1][STATUS[j]]
			if prs {
				if prev.aCount <= 1 {
					cache[i]['A'] = &History{prev.aCount + 1, prev.endingLCount}
					count = (count + 1) % MOD
				}
				if prev.endingLCount <= 2 {
					cache[i]['L'] = &History{prev.aCount, prev.endingLCount + 1}
					count = (count + 1) % MOD
				}
				cache[i]['P'] = &History{prev.aCount, prev.endingLCount}
				count = (count + 1) % MOD
			}
		}
	}
	return count
}
