package regular

import "fmt"

//func isPowerOfTwo(n int) bool {
//	if n == 0 {
//		return false
//	}
//	for n != 1 {
//		if n&0x1 == 1 {
//			return false
//		}
//		n = n >> 1
//	}
//	return true
//}
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1) == 0)
}

func T231() {
	fmt.Println(isPowerOfTwo(3))
}
