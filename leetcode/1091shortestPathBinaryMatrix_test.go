package leetcode

import (
	"fmt"
	"testing"
)

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	n := len(grid)
	if grid[n-1][n-1] == 1 {
		return -1
	}
	ans := 0
	// breadth first search, 声明队列，初始起点0，0入队，并将起点置为visited,1
	grid[0][0] = 1
	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		ans++
		for k := len(queue); k > 0; k-- {
			// 出队
			s := queue[0]
			// 位置
			i, j := s[0], s[1]
			// 剩余队列
			queue = queue[1:]
			// 判断是否是目标点
			if i == n-1 && j == n-1 {
				return ans
			}
			// i左右减1的范围是下个候补点
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 0 {
						//设置为visited
						grid[x][y] = 1
						queue = append(queue, [2]int{x, y})
					}
				}
			}
		}
	}
	return -1

}
func TestShortestPathBinaryMatrix(t *testing.T) {
	grip := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grip))

}
