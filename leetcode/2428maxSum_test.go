package leetcode

import (
	"fmt"
	"testing"
)

func maxSum(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return -1
	}
	for {

	}

	return -1
}

func TestMaxSum(t *testing.T) {
	/**
	1,2,3,5
	4,5,6,6
	7,8,9,9
	*/
	grip := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(maxSum(grip))

}
