package regular

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor297() Codec {
	codec := Codec{}
	return codec
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	parents := []int{}
	values := []int{}
	var iterate func(node *TreeNode, parent int)
	iterate = func(node *TreeNode, parent int) {
		if node == nil {
			values = append(values, 10001)
			parents = append(parents, parent)
			return
		}
		cur := len(parents)
		values = append(values, node.Val)
		parents = append(parents, parent)
		if node.Left != nil || node.Right != nil {
			iterate(node.Left, cur)
			iterate(node.Right, cur)
		}
	}
	iterate(root, -1)
	result := ""
	for i, parent := range parents {
		result += strconv.Itoa(parent)
		if i != len(parents)-1 {
			result += ","
		}
	}
	result += "|"
	for i, value := range values {
		result += strconv.Itoa(value)
		if i != len(values)-1 {
			result += ","
		}
	}
	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	parents := []int{}
	values := []int{}
	indexAndValue := strings.Split(data, "|")
	indices := strings.Split(indexAndValue[0], ",")
	vals := strings.Split(indexAndValue[1], ",")
	for _, index := range indices {
		iIndex, _ := strconv.Atoi(index)
		parents = append(parents, iIndex)
	}
	for _, value := range vals {
		iValue, _ := strconv.Atoi(value)
		values = append(values, iValue)
	}
	parentMap := make(map[int][]int)
	for i, parent := range parents {
		val, exist := parentMap[parent]
		if !exist {
			val = []int{}
		}
		val = append(val, i)
		parentMap[parent] = val
	}

	root := &TreeNode{}
	var compose func(int, *TreeNode)
	compose = func(index int, parentNode *TreeNode) {
		children := parentMap[index]
		if children != nil {
			if values[children[0]] != 10001 {
				parentNode.Left = &TreeNode{children[0], nil, nil}
				compose(children[0], parentNode.Left)
			}
			if values[children[1]] != 10001 {
				parentNode.Right = &TreeNode{children[1], nil, nil}
				compose(children[1], parentNode.Right)
			}
		}
	}
	compose(0, root)
	var iterate func(node *TreeNode)
	iterate = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Val = values[node.Val]
		if node.Left != nil {
			iterate(node.Left)
		}
		if node.Right != nil {
			iterate(node.Right)
		}
	}
	iterate(root)
	return root
}

func T297() {
	codec := Constructor297()
	//tree := &TreeNode{10, &TreeNode{20, &TreeNode{40, &TreeNode{60, nil, nil}, nil}, nil}, &TreeNode{30, nil, &TreeNode{50, nil, nil}}}
	tree := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	sed := codec.serialize(tree)
	fmt.Println(sed)
	tree = codec.deserialize(sed)
	fmt.Println(tree)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor5903();
 * deser := Constructor5903();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
