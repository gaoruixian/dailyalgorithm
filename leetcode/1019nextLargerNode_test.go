package leetcode

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	var stack []int
	res := make([]int, 0)
	i := 0
	for head != nil {
		res = append(res, 0)
		for len(stack) > 0 && res[stack[len(stack)-1]] < head.Val {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = head.Val
		}
		stack = append(stack, i)
		i++
		head = head.Next
	}
	return res

}

func nextLargerNodes2(head *ListNode) []int {
	/**
	2,1,5
	*/
	res, indexes, nums := make([]int, 0), make([]int, 0), make([]int, 0)
	p := head
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}

	for i := 0; i < len(nums); i++ {
		res = append(res, 0)
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for len(indexes) > 0 && nums[indexes[len(indexes)-1]] < num {
			index := indexes[len(indexes)-1]
			res[index] = num
			indexes = indexes[:len(indexes)-1]
			fmt.Println("inner", indexes, i)
		}
		indexes = append(indexes, i)
		fmt.Println("indexes", indexes)
	}
	return res

}

func TestNext(t *testing.T) {
	node := new(ListNode)
	node.Val = 2
	/**
	[2,1,5]
	*/
	node.Next = &ListNode{1, &ListNode{5, nil}}
	//fmt.Println(nextLargerNodes(node))
	fmt.Println(nextLargerNodes2(node))

}
