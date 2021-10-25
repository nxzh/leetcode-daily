package y2021m08

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	levels := []*Node{root}
	var result [][]int
	for len(levels) > 0 {
		size := len(levels)
		var elemsOfLevel []int
		for i := 0; i < size; i++ {
			elem := levels[i]
			elemsOfLevel = append(elemsOfLevel, elem.Val)
			if elem.Children != nil {
				for j := 0; j < len(elem.Children); j++ {
					if elem.Children[j] != nil {
						levels = append(levels, elem.Children[j])
					}
				}
			}
		}
		levels = levels[size:]
		result = append(result, elemsOfLevel)
	}
	return result
}
