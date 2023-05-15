package leetcode

import (
	"fmt"
	"testing"
)
/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
 */
func TestTwoSum(t *testing.T)  {
	fmt.Println(twoSum([]int{2,5,6,7},9))
	fmt.Println(twoSum2([]int{2,5,6,7},9))
}

func twoSum(nums []int, target int) []int  {
	if nums == nil || len(nums) < 2{
		return nil
	}
	for i,v := range nums {
		other := target - v
		for j:= i+1;j<len(nums);j++  {
			if nums[j] == other {
				return []int{i,j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int ,target int) []int   {
	if nums == nil || len(nums) < 2 {
		return nil
	}
	m := make(map[int]int)
	for index,value := range nums{
		find := target - value
		if v,ok := m[find];ok {
			fmt.Println(v,ok)
			return []int{index,v}
		}else {
			m[value] = index
		}
	}
	fmt.Println(m)
	return nil
}