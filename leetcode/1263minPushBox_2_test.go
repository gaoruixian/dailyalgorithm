package leetcode

import (
	"fmt"
	"testing"
)

/*
* 判断这个（x,y）位置是否能停留
 */
func isCanStay(grip [][]byte, m, n, x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n && grip[x][y] != '#'
}

// 上下左右
var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// startX,startY --> targetX,targetY
func personCanReach(grip [][]byte, visited [][]bool, m, n, startX, startY, targetX, targetY int) bool {
	if startX == targetX && startY == targetY {
		return true
	}
	visited[startX][startY] = true
	for _, d := range directions {
		nextX := startX + d[0]
		nextY := startY + d[1]
		if isCanStay(grip, m, n, nextX, nextY) && !visited[nextX][nextY] {
			if personCanReach(grip, visited, m, n, nextX, nextY, targetX, targetY) {
				//fmt.Println("people", startX, startY, targetX, targetY)
				return true
			}
		}
	}
	return false
}

type Box struct {
	x    int
	y    int
	from int
}

func minPushBox2(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	var bx, by, sx, sy, targetStartX, targetStartY int
	// 遍历找到初始位置
	for i, l := range grid {
		for j, v := range l {
			if v == 'B' {
				bx = i
				by = j
			}
			if v == 'S' {
				sx = i
				sy = j
			}
			if v == 'T' {
				targetStartX = i
				targetStartY = j
			}
		}
	}
	//fmt.Println(bx, by, sx, sy, targetStartX, targetStartY)
	//初始化队列，加入元素以启动BFS（breadth-first search）
	var visitedBefore = make([][][4]bool, m)
	for i := range visitedBefore {
		visitedBefore[i] = make([][4]bool, n)
	}
	var visited = make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	//visited[m][n] = true
	queue := make([]Box, 0)
	for i := 0; i < 4; i++ {
		direction := directions[i]
		if personCanReach(grid, visited, m, n, sx, sy, bx+direction[0], by+direction[1]) {
			queue = append(queue, Box{bx, by, i})
			//fmt.Println(bx, by, i)
			visitedBefore[bx][by][i] = true
		}
	}
	// 以箱子为视角开始BFS
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			grid[cur.x][cur.y] = 'B'
			sx = cur.x - directions[cur.from][0]
			sy = cur.y - directions[cur.from][1]
			fmt.Println(cur)
			if cur.x == targetStartX && cur.y == targetStartY {
				return step
			}
			for j := 0; j < 4; j++ {
				direction := directions[j]
				nextX := cur.x + direction[0]
				nextY := cur.y + direction[1]
				//var visited [][]bool = [m][n]
				// 人是否能绕到箱子的后面？
				if !personCanReach(grid, visited, m, n, sx, sy, cur.x-direction[0], cur.y-direction[1]) {
					continue
				}
				// 箱子的下个位置是否合法？
				if !isCanStay(grid, m, n, nextX, nextY) {
					continue
				}
				// 箱子的下一个状态是不是重复了？
				if visitedBefore[nextX][nextY][i] {
					continue
				}
				queue = append(queue, Box{nextX, nextY, i})
				visitedBefore[nextX][nextY][i] = true
			}
			grid[cur.x][cur.y] = '.'
		}
		step++
	}
	return -1
}

func TestPushBox2(t *testing.T) {
	grid := [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '.', '.', '#', '#'},
		{'#', '.', '#', 'B', '.', '#'},
		{'#', '.', '.', '.', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'},
	}
	// 向下、向左、向左、向上再向上。
	fmt.Println(minPushBox2(grid))
	grid = [][]byte{{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '.', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'}}
	//fmt.Println(minPushBox2(grid))
}
