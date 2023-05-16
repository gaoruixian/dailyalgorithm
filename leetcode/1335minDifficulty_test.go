package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// 1 <= jobDifficulty.length <= 300
// 0 <= jobDifficulty[i] <= 1000
// 1 <= d <= 10
func minDifficulty2(jobDifficulty []int, d int) int {
	size := len(jobDifficulty)
	if d > size {
		return -1
	}
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, d)
	}
	// dp[i][j] i项工作，j天完成最低难度 dp[i-1][j]+max
	return dp[size-1][d-1]

}
func minDifficulty3(jobDifficulty []int, d int) int {
	size := len(jobDifficulty)
	if d > size {
		return -1
	}
	dp := make([]int, size)
	for i, j := 0, 0; i < size; i++ {
		j := max1(j, jobDifficulty[i])
		dp[i] = j
	}
	for i := 1; i < d; i++ {
		ndp := make([]int, size)
		for index := 0; index < size; index++ {
			ndp[index] = 1001
		}
		for j := i; j < size; j++ {
			ma := 0
			for k := j; k >= i; k-- {
				ma = max1(ma, jobDifficulty[k])
				ndp[j] = min(ndp[j], ndp[k-1]+ma)
			}
		}
		dp = ndp
	}
	return dp[size-1]

}

func minDifficulty(jobDifficulty []int, d int) int {
	jobLen := len(jobDifficulty)
	if len(jobDifficulty) < d {
		return -1
	}
	dp := make([][]int, jobLen+1) // dp[nJob][nDay]表示前nJob个任务分配到nDay天的困难度
	dp[0] = make([]int, d+1)
	for i := 0; i < d+1; i++ {
		dp[0][i] = 300001
	}
	dp[0][0] = 0

	for nJob := 1; nJob <= jobLen; nJob++ {
		dp[nJob] = make([]int, d+1)
		dp[nJob][0] = 300001

		//因为每天至少要有一个任务，所以nJob个任务不可能分配到nJob+1以上的天中
		for nDay := 1; nDay <= nJob && nDay < d+1; nDay++ {
			md := 0
			dp[nJob][nDay] = 300001
			//这里 从后往前，就可以快速获取中当前天的最大值
			for j := nJob - 1; j >= nDay-1; j-- {
				if jobDifficulty[j] > md {
					md = jobDifficulty[j]
				}
				//比较 前j个任务分配到nDay-1天的困难程度 加上 剩余任务（分配到了nday天）的最大困难程度
				if dp[j][nDay-1]+md < dp[nJob][nDay] {
					dp[nJob][nDay] = dp[j][nDay-1] + md
				}
			}
		}
	}
	return dp[jobLen][d]
}

func TestMinDifficulty(t *testing.T) {
	jobDifficulty := []int{11, 111, 22, 222, 33, 333, 44, 444}
	day := 6
	fmt.Println(minDifficulty(jobDifficulty, day))

}
func TestA(t *testing.T) {
	fmt.Println(math.MaxInt32 / 2)

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxOfArray(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
