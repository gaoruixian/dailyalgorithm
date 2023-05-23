package leetcode

import (
	"fmt"
	"testing"
)

func numTilePossibilities(tiles string) int {
	size := len(tiles)
	countMap := make(map[string]int)
	for _, v := range tiles {
		countMap[string(v)]++
	}
	//fmt.Println(countMap)
	// A，A，B
	for i := 1; i <= size; i++ {

	}
	return dfs(len(tiles), countMap) - 1

}
func dfs(i int, count map[string]int) int {
	fmt.Println(i, ",", count)
	if i == 0 {
		return 1
	}
	//fmt.Println(i, ",", count)
	res := 1
	for t := range count {
		if count[t] > 0 {
			count[t]--
			fmt.Println(t, "before", count)
			res = res + dfs(i-1, count)
			count[t]++
			fmt.Println(t, "after:", count)
		}
	}
	return res
}
func TestNumTitlePossibilities(t *testing.T) {
	tiles := "AAB"
	//tiles = "AAABBC"

	fmt.Println(numTilePossibilities(tiles))
}

func Test1(t *testing.T) {
	a := "A"
	tiles := "AAB"
	tiles = "AAABBC"
	for _, v := range tiles {
		fmt.Println(string(v))

	}
	fmt.Println(a)
}
