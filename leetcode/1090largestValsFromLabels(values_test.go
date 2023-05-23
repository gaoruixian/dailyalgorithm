package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/*
values := []int{5, 4, 3, 2, 1}
labels := []int{1, 1, 2, 2, 3}
numWanted := 3
useLimit := 1
*/
type compareInt []int

func (a compareInt) Len() int {
	return len(a)
}

func (a compareInt) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a compareInt) Less(i, j int) bool {
	return a[i] > a[j]
}
func largestValsFromLabels2(values []int, labels []int, numWanted int, useLimit int) int {
	ans := 0
	n := len(values)
	// 下表数组，将values和labels数组联系起来
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	// 降序
	sort.Slice(idx, func(i, j int) bool {
		return values[idx[i]] > values[idx[j]]
	})
	// label计数
	cnt := make(map[int]int)
	used := 0
	for i := 0; i < n; i++ {
		label := labels[idx[i]]
		if cnt[label] >= useLimit {
			continue
		}
		used++
		ans += values[idx[i]]
		cnt[label]++
		if used == numWanted {
			break
		}
	}
	return ans
}
func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	ans := 0
	n := len(values)
	for i := 1; i <= numWanted; i++ {
		labelsMap := make(map[int]int)
		for j := 0; j < n; j++ {

			for k := j + 1; k <= j; k++ {
				if labelsMap[labels[j]] > useLimit {
					continue
				}
				labelsMap[labels[j]]++
				ans = ans + values[k]

			}

		}
		ans = ans + values[i]

	}
	return ans

}
func TestLargestValsFromLabels(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	labels := []int{1, 1, 2, 2, 3}
	numWanted := 3
	useLimit := 1
	fmt.Println(largestValsFromLabels(values, labels, numWanted, useLimit))
}
func TestLargest(t *testing.T) {
	values := []int{5, 4, 3, 1, 2}
	n := len(values)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	fmt.Println(idx)
	//sort.Slice(idx, func(i, j int) bool {
	//	return values[idx[i]] > values[idx[j]]
	//})
	fmt.Println(idx)
	fmt.Println(values)
	sort.Sort(compareInt(idx))
	fmt.Println(idx)

}
