package y2021m08

import (
	"fmt"
	"math"
	"sort"
)

func rectangleArea(rectangles [][]int) int {
	if len(rectangles) == 1 {
		return (rectangles[0][2] - rectangles[0][0]) * (rectangles[0][3] - rectangles[0][1]) % 1000000007
	}
	sort.Slice(rectangles, func(i, j int) bool {
		if rectangles[i][0] < rectangles[j][0] {
			return true
		} else if rectangles[i][0]-rectangles[j][0] == 0 {
			return rectangles[i][1]-rectangles[j][1] <= 0
		} else {
			return false
		}
	})
	cur := rectangles[0]
	var mCur [][]bool
	total := 0
	curSquare := (cur[2] - cur[0]) * (cur[3] - cur[1])
	for i := 1; i < len(rectangles); i++ {
		newR := rectangles[i]
		if mCur == nil {
			mCur = matrix(cur, []int{0, 0}, cur[2]-cur[0], cur[3]-cur[1])
		}
		if insect(cur, rectangles[i]) {
			newCur := []int{
				int(math.Min(float64(cur[0]), float64(newR[0]))),
				int(math.Min(float64(cur[1]), float64(newR[1]))),
				int(math.Max(float64(cur[2]), float64(newR[2]))),
				int(math.Max(float64(cur[3]), float64(newR[3]))),
			}
			l := newCur[2] - newCur[0]
			w := newCur[3] - newCur[1]
			mCur = expandMatrix(mCur, []int{cur[0] - newCur[0], cur[1] - newCur[1]}, l, w)
			mNew := matrix(newR, []int{newR[0] - newCur[0], newR[1] - newCur[1]}, l, w)
			mCur, curSquare = countInsect(mCur, mNew)
			cur = newCur
		} else {
			total += curSquare
			curSquare = (newR[2] - newR[0]) * (newR[3] - newR[1])
			cur = newR
			mCur = nil
		}
	}
	total += curSquare
	return total % 1000000007
}
func countInsect(m1 [][]bool, m2 [][]bool) ([][]bool, int) {
	l := len(m1)
	w := len(m1[0])
	c := 0
	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			m1[i][j] = m1[i][j] || m2[i][j]
			if m1[i][j] {
				c++
			}
		}
	}
	return m1, c
}
func insect(rect1 []int, rect2 []int) bool {
	if rect2[0] < rect1[2] {
		return true
	}
	return false
}

func expandMatrix(m [][]bool, pos []int, l int, w int) [][]bool {
	matrix := make([][]bool, l)
	for i := 0; i < l; i++ {
		matrix[i] = make([]bool, w)
	}
	posX, posY := pos[0], pos[1]
	rectL, rectW := len(m), len(m[0])
	for i := posX; i < posX+rectL; i++ {
		for j := posY; j < posY+rectW; j++ {
			matrix[i][j] = m[i-posX][j-posY]
		}
	}
	return matrix
}

func matrix(rect []int, pos []int, l int, w int) [][]bool {
	matrix := make([][]bool, l)
	for i := 0; i < l; i++ {
		matrix[i] = make([]bool, w)
	}
	posX, posY := pos[0], pos[1]
	rectL, rectW := rect[2]-rect[0], rect[3]-rect[1]
	for i := posX; i < posX+rectL; i++ {
		for j := posY + rectW - 1; j >= posY; j-- {
			matrix[i][j] = true
		}
	}
	return matrix
}

func Call0822() {
	//res := rectangleArea([][]int{{0, 0, 1, 1}, {2, 2, 3, 3}}) //2
	//[[25,20,70,27],[68,80,79,100],[37,41,66,76]]
	res := rectangleArea([][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}, {3, 0, 7, 2}, {5, 2, 8, 3}})
	//res := rectangleArea([][]int{{39,99,60,100},{69,14,91,56},{13,42,20,70}}) //2
	fmt.Println(res)
}
