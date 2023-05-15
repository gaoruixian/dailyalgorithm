package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func queryString(s string, n int) bool {
	sum := 1
	if n == 1 {
		sum = 1
	} else {
		sum = (1 + n) * n / 2
	}
	size := len(s)
	binarySum := 0
	for i, v := range s {
		if v == '1' {
			binarySum += 1 << (size - 1 - i)
		}
	}
	fmt.Println(sum, binarySum)
	if sum > binarySum {
		return false
	} else {
		return true
	}
}

func queryString2(s string, n int) bool {
	for i := 1; i <= n; i++ {
		if !strings.Contains(s, strconv.FormatInt(int64(i), 2)) {
			return false
		}
	}
	return true
}
func TestQueryString(t *testing.T) {
	s := "10010111100001110010"
	n := 10
	fmt.Println(queryString(s, n))
	fmt.Println(queryString2(s, n))

}
