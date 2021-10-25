package y2021m09

import (
	"math"
	"sort"
)

func outerTrees(trees [][]int) [][]int {
	sort.Slice(trees, func(i, j int) bool {
		return trees[i][0] < trees[j][0]
	})
	fence := [][]int{}
	fenceHigh := append(fence, trees[0])
	slopeHigh := []int{math.MaxInt32}
	fenceLow := append(fence, trees[0])
	slopeLow := []int{math.MinInt32}
	for i := 1; i < len(trees); i++ {
		tree := trees[i]
		curHigh := fenceHigh[len(fenceHigh)-1]
		curLow := fenceLow[len(fenceLow)-1]
		slopeWithHigh := slope(curHigh, tree)
		slopeWithLow := slope(curLow, tree)
		if slopeWithHigh <= slopeHigh[len(slopeHigh)-1] {
			slopeHigh = append(slopeHigh, slopeWithHigh)
			fenceHigh = append(fenceHigh, tree)
		} else {
		}
		if slopeWithLow >= slopeLow[len(slopeLow)-1] {
			slopeLow = append(slopeLow, slopeWithLow)
			fenceLow = append(fenceLow, tree)
		} else {

		}
	}
	return nil
}
func findPos(slopes []int, slope float64) int {

	return 0
}
func slope(pt1 []int, pt2 []int) int {
	if pt2[1] == pt1[1] {
		return 0
	}
	return (pt2[1] - pt1[1]) / (pt2[0] - pt1[0])
}
