package y2021m09

func sumOfDistancesInTree(n int, edges [][]int) []int {
	dist := make(map[int]map[int]int)

	for i := 0; i < len(edges); i++ {
		if _, prs := dist[edges[i][0]]; !prs {
			dist[edges[i][0]] = make(map[int]int)
		}
		dist[edges[i][0]][edges[i][1]] = 1
		if _, prs := dist[edges[i][1]]; !prs {
			dist[edges[i][1]] = make(map[int]int)
		}
		dist[edges[i][1]][edges[i][0]] = 1
	}
	return solution(n, dist)
}

func solution(n int, dist map[int]map[int]int) []int {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			sum += calcDist(i, j, i, 0, dist)
		}
	}
	return nil
}

func calcDist(from int, to int, origin int, cur int, dist map[int]map[int]int) int {

	return 0
}

func Call04() {

}
