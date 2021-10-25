package y2021m09

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var key = keysPressed[0]
	var time = releaseTimes[0]
	for i := 1; i < len(keysPressed); i++ {
		cur := releaseTimes[i] - releaseTimes[i-1]
		if cur > time {
			key = keysPressed[i]
			time = cur
		} else if cur == time {
			if keysPressed[i] > key {
				key = keysPressed[i]
			}
		}
	}
	return key
}
