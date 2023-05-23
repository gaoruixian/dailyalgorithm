package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset2(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	limit = limit - root.Val
	// 叶子节点
	if root.Left == root.Right {
		if limit > 0 {
			return nil
		}
		return root
	}
	// 左子节点
	root.Left = sufficientSubset2(root.Left, limit)
	// 右子节点
	root.Right = sufficientSubset2(root.Right, limit)
	// 左右子节点都为空，删除该节点
	if root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	type FuncType func(node *TreeNode, pre, limit int) int
	var dfs FuncType
	dfs = func(node *TreeNode, pre, limit int) int {
		if node == nil {
			return math.MinInt
		}
		// 叶子节点
		if node.Left == nil && node.Right == nil {
			return node.Val + pre
		}
		left := dfs(node.Left, node.Val+pre, limit)
		right := dfs(node.Right, node.Val+pre, limit)
		if left < limit {
			node.Left = nil
		}
		if right < limit {
			node.Right = nil
		}
		if left > right {
			return left
		} else {
			return right
		}
	}

	if dfs(root, 0, limit) < limit {
		return nil
	}
	return root
}
func TestSufficientSubset(t *testing.T) {
	/*
		输入：[5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
	*/
	root := &TreeNode{
		5,
		&TreeNode{4,
			&TreeNode{11,
				&TreeNode{7, nil, nil},
				&TreeNode{1, nil, nil}},
			nil},
		&TreeNode{8,
			&TreeNode{17, nil, nil},
			&TreeNode{4,
				&TreeNode{5, nil, nil},
				&TreeNode{3, nil, nil}}}}
	limit := 22
	//fmt.Println(*root)
	root = sufficientSubset(root, limit)
	traverse(root)

}
func traverse(node *TreeNode) {
	if node == nil {
		return
	}

	// 递归遍历左子树
	traverse(node.Left)

	// 输出当前节点的值
	fmt.Println(node.Val)

	// 递归遍历右子树
	traverse(node.Right)
}
func TestLeaf(t *testing.T) {
	root := &TreeNode{3, nil, nil}
	if root.Left == root.Right {
		fmt.Println("same")
	}

}
