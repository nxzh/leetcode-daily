package y2021m08

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	nums1 := strings.Split(num1, "+")
	nums2 := strings.Split(num2, "+")
	r1, _ := strconv.Atoi(nums1[0])
	v1, _ := strconv.Atoi(nums1[1][:len(nums1[1])-1])
	r2, _ := strconv.Atoi(nums2[0])
	v2, _ := strconv.Atoi(nums2[1][:len(nums2[1])-1])

	return strconv.Itoa(r1*r2-v1*v2) + "+" + strconv.Itoa(r1*v2+r2*v1) + "i"

}

func Call0824() {
	res := complexNumberMultiply("1+-1i", "1+-1i")
	fmt.Println(res)
}
