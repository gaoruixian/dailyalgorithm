package leetcode

import (
	"fmt"
	"testing"
)

// 1 <= arr1.length, arr2.length <= 1000
// arr1[i] 和 arr2[i] 都是 0 或 1
// arr1 和 arr2 都没有前导0
func addNegabinary(arr1 []int, arr2 []int) []int {
	size := max(len(arr1), len(arr2))
	ans := make([]int, size+1)
	a, b := 0, 0

	for i, v := range arr1 {
		if i == 0 {
			a += v * 1
		} else if v == 1 && i%2 == 0 {
			a += 2 << (i - 1)
		} else if v == 1 && i%2 != 0 {
			a += 2 << (i - 1) * (-1)
		}
	}
	fmt.Println(a)
	for i, v := range arr2 {
		if i == 0 {
			b += v * 1
		} else if v == 1 && i%2 == 0 {
			b += 2 << (i - 1)
		} else if v == 1 && i%2 != 0 {
			b += 2 << (i - 1) * (-1)
		}
	}
	fmt.Println(b)

	return ans
}
func TestNegativeBinary(t *testing.T) {
	arr1 := []int{1, 1, 1, 1, 1}
	arr2 := []int{1, 0, 1}
	fmt.Println(addNegabinary(arr1, arr2))

}
