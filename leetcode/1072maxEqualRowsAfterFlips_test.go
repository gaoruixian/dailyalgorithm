package leetcode

import (
	"fmt"
	"testing"
)

// 给定 m x n 矩阵 matrix
func maxEqualRowsAfterFlips(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	cnt := make([]int, m)
	for row, rowV := range matrix {
		for _, v := range rowV {
			if v == 1 {
				cnt[row]++
			}
		}
	}
	fmt.Println(cnt)
	rowSum := 0
	rowCount := 1
	for _, v := range cnt {
		rowSum += v
		if rowSum == 0 || rowSum%n == 0 {
			rowCount++
		}
	}
	return rowCount
}

func maxEqualRowsAfterFlips2(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	rowCount := 1
	for i := 0; i < m-1; i++ {
		rowSum := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[i+1][j] {
				rowSum += 1
			}
		}
		if rowSum%n == 0 {
			rowCount++
		}
	}
	return rowCount
}
func TestMaxEqualRowsAfterFlips(t *testing.T) {
	matrix := [][]int{
		{1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
		{1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1}}
	//fmt.Println(maxEqualRowsAfterFlips(matrix))
	//matrix = [][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}
	fmt.Println(maxEqualRowsAfterFlips(matrix))
	fmt.Println(maxEqualRowsAfterFlips2(matrix))

}
