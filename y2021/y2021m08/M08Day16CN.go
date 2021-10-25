package y2021m08

func countArrangement(n int) int {
	return len(count(n))
}

func count(n int) [][]byte {
	if n == 1 {
		return [][]byte{{1}}
	} else {
		prev := count(n - 1)
		for i := 0; i < len(prev); i++ {
			prev[i] = append(prev[i], byte(n))
		}
		cache := make(map[string]bool)
		for i := 0; i < len(prev); i++ {
			ary := prev[i]
			if !cache[string(ary)] {
				lenAry := len(ary)
				cache[string(ary)] = true
				for j := 0; j < lenAry-1; j++ {
					if (n%(j+1) == 0 || (j+1)%n == 0) && (int(ary[j])%n == 0 || n%int(ary[j]) == 0) {
						copied := make([]byte, len(ary))
						copy(copied, ary)
						temp := copied[j]
						copied[j] = copied[lenAry-1]
						copied[lenAry-1] = temp
						if !cache[string(copied)] {
							prev = append(prev, copied)
							cache[string(copied)] = true
						}
					}
				}
			}
		}
		return prev
	}
}
