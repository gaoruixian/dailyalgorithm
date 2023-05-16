package dynamicprogramming

import (
	"fmt"
	"testing"
)

/*
*
一个经典的动态规划问题是求解斐波那契数列。它的定义是：F(0) = 0, F(1) = 1, F(n) = F(n-1) + F(n-2) (n >= 2)。
*/
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
func TestFibonacci(t *testing.T) {
	n := 4
	fmt.Println(fibonacci(n))
}

/*
*
C 包内总容量
n 物品数量
w 每个物品占用容量
v 每个物品价值
求包内能放入物品最大价值
*/
func knapsack2(C int, n int, w []int, v []int) int {
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, C+1)
	}
	/**
	设f(i, j)表示前i个物品放入容量为j的背包中所能获得的最大价值。
	f[i][j]前i个物品放入容量为j的背包中获得的最大价值
	*/

	for i := 1; i <= n; i++ {
		for j := 1; j <= C; j++ {
			// 物品容量大于包
			if j < w[i] {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = max(f[i-1][j], f[i-1][j-w[i]]+v[i])
			}
		}
	}
	return f[n][C]

}
func knapsack(C int, n int, w []int, v []int) int {
	// 创建一个二维切片f来存储每个状态的最大价值，初始化为0
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, C+1)
	}
	/**
	设f(i, j)表示前i个物品放入容量为j的背包中所能获得的最大价值。
	f[i][j]前i个物品放入容量为j的背包中获得的最大价值
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= C; j++ {
			// 物品容量大于包,包内放不下第i个物体
			if j < w[i-1] {
				// 数组下标从0开始，在进行状态转移时，需要将i和j分别减去1
				f[i][j] = f[i-1][j]
				// 包容量大于物品，可以放入第i个物品。此时比较放入物体和不放入物体的价值
			} else {
				f[i][j] = max(f[i-1][j], f[i-1][j-w[i-1]]+v[i-1])
			}
		}
	}
	return f[n][C]
}

func max(a, b int) int {
	if a > b {
		return a

	}
	return b
}
func TestKnapsack(t *testing.T) {
	C := 20
	n := 11
	w := []int{2, 5, 6, 3, 1, 8, 7, 5, 1, 1, 1}
	v := []int{1, 3, 4, 3, 1, 6, 4, 5, 1, 1, 1}
	fmt.Println(knapsack(C, n, w, v))
	//fmt.Println(knapsack2(C, n, w, v))

}
