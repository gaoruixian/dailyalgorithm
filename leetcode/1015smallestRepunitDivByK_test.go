package leetcode

import (
	"fmt"
	"testing"
)

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 {
		return -1
	}

	return -1
}

func TestK(t *testing.T) {
	k := 4
	fmt.Println(smallestRepunitDivByK(k))

}
