package leetcode

import (
	"fmt"
	"testing"
)

// 给定 m x n 矩阵 matrix
func maxEqualRowsAfterFlips(matrix [][]int) int {
	for row, rowV := range matrix {
		for col, v := range rowV {

		}
	}
	return 1

}
func TestMaxEqualRowsAfterFlips(t *testing.T) {
	matrix := [][]int{{0, 1, 1}, {1, 0, 0}, {0, 0, 1}}
	fmt.Println(maxEqualRowsAfterFlips(matrix))

}
