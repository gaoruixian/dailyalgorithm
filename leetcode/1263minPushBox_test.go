package leetcode

import (
	"fmt"
	"testing"
)

/**
首先需要找出广度优先搜索中我们需要一直向外层扩张的状态结构。虽然本题要求的是最少的推箱子次数，但是因为需要人来推动箱子，
就需要先判断人所在的状态能否满足我们想要让箱子移动的状态。所以我们bfs的队列中应该同时保存箱子和人的状态信息。
除了基本的箱子移动方向的bfs之外，还需要一次判断人能否到达推箱子位置dfs和箱子能否被推到目标位置的dfs，共计3次。
*/
// 定义结构体
type Node struct {
	Px int
	Py int // 人的坐标
	Bx int
	By int // 箱子坐标
}

// 人能否到达推箱子的位置
// cur表示当前的箱子和人的状态，tx,ty 表示人要到达的推箱子的位置
func validPosition(cur *Node, grid [][]byte, tx, ty int) bool {
	m, n := len(grid), len(grid[0])
	// 保存已经处理过的位置
	tried := make(map[string]bool)
	var dfs func(int, int) bool
	dfs = func(x int, y int) bool {
		// 如果现在处理的位置不可以到达，或者已经处理过了，返回false
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '#' || tried[fmt.Sprintf("%d%d", x, y)] {
			return false
		}
		// 如果当前位置是箱子的位置，同样不可以到达
		if x == cur.Bx && y == cur.By {
			return false
		}
		// 如果到达了目标位置，返回true
		if x == tx && y == ty {
			return true
		}
		// 标记当前位置已经处理过
		tried[fmt.Sprintf("%d%d", x, y)] = true
		// 处理当前坐标的相邻四个位置
		return dfs(x+1, y) || dfs(x-1, y) || dfs(x, y+1) || dfs(x, y-1)
	}
	// 从人当前的位置开始处理
	return dfs(cur.Px, cur.Py)
}

// 箱子能否移动到 bx+dx,by+dy 坐标位置
// cur表示当前箱子和人的状态，dx，dy分别表示箱子的偏移距离
// next是箱子移动后箱子和人的状态，如果不能移动则返回nil
func validForward(cur *Node, grid [][]byte, dx, dy int) (next *Node) {
	// 当前箱子坐标+偏移量，计算出箱子目标坐标
	bx := cur.Bx + dx
	by := cur.By + dy
	m, n := len(grid), len(grid[0])
	// 如果目标坐标不可以到达，直接返回nil
	if bx < 0 || bx >= m || by < 0 || by >= n || grid[bx][by] == '#' {
		return
	}
	//  bx-dx,by-dy 是将箱子推向目标位置时，人的位置
	//  如果人不能到达该位置，返回nil
	if !validPosition(cur, grid, cur.Bx-dx, cur.By-dy) {
		return
	}
	// 如果上面的判断都满足，表示可以移动箱子到指定便宜位置
	// 返回next，人的位置就是cur的箱子位置，箱子位置就是cur箱子位置+偏移量
	next = &Node{Px: cur.By, Py: cur.By, Bx: bx, By: by}
	return
}

func minPushBox(grid [][]byte) int {
	start, target := &Node{}, &Node{}
	// 处理start和target状态，target中人的状态还不能确定
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				start.Px = i
				start.Py = j
				grid[i][j] = '.'
			}
			if grid[i][j] == 'B' {
				start.Bx = i
				start.By = j
			}
			if grid[i][j] == 'T' {
				target.Bx = i
				target.By = j
				grid[i][j] = '.'
			}
		}
	}
	// 偏移量的辅助数组
	dirs := []int{0, -1, 0, 1, 0}
	// 记录已经处理过的状态，这个状态包含了箱子和人的状态
	visited := make(map[string]bool)
	// queue中保存的我们自定义的状态node
	queue := make([]*Node, 0)
	// 从start状态开始bfs
	queue = append(queue, start)
	//  将start标记为访问过
	visited[fmt.Sprintf("%d%d%d%d", start.Px, start.Py, start.Bx, start.By)] = true
	//  steps是bfs搜索的层数，也就是推箱子的次数
	steps := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 向cur的四个相邻方向开始bfs
			for j := 0; j < 4; j++ {
				// 判断是否能到达相邻方向，拿到next状态
				next := validForward(cur, grid, dirs[j], dirs[j+1])
				if next == nil {
					continue
				}
				// 标记一个已经访问过的状态
				key := fmt.Sprintf("%d%d%d%d", next.Px, next.Py, next.Bx, next.By)
				if visited[key] {
					continue
				}
				// 如果要移动到的下一个箱子的位置就是终点，答案就是step+1
				// bfs保证了第一个答案就是最短路径
				if next.Bx == target.Bx && next.By == target.By {
					return steps + 1
				}
				visited[key] = true
				// 将一个合法的next状态加入队列
				queue = append(queue, next)
			}
		}
		steps++
	}
	return -1
}

func TestPushBox(t *testing.T) {
	grip := [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '.', '.', '#', '#'},
		{'#', '.', '#', 'B', '.', '#'},
		{'#', '.', '.', '.', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'},
	}
	// 向下、向左、向左、向上再向上。
	fmt.Println(minPushBox(grip))

}
