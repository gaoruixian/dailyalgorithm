package leetcode

import (
	"fmt"
	"testing"
)

/*
*
croak
*/
func minNumberOfFrogs(croakOfFrogs string) int {
	if len(croakOfFrogs)%5 != 0 {
		return -1
	}
	frogMap := map[int32]int{
		'c': 0,
		'r': 1,
		'o': 2,
		'a': 3,
		'k': 4,
	}
	var cnt = [5]int{}
	var newFrog = 0
	var ans = 0
	for _, v := range croakOfFrogs {
		index := frogMap[v]
		// new frog
		if index == 0 {
			newFrog++
			cnt[0]++
			if newFrog > ans {
				ans = newFrog
			}
			// old frog croak
		} else {
			if cnt[index-1] == 0 {
				return -1
			}
			cnt[index-1]--
			if index == 4 {
				newFrog--
			} else {
				cnt[index]++
			}
		}
	}
	if newFrog > 0 {
		return -1
	} else {
		return ans
	}

}

// 每个字母在 "croak" 中的上一个字母
var previous = [...]int{'c': 'k', 'r': 'c', 'o': 'r', 'a': 'o', 'k': 'a'}

func minNumberOfFrogs2(croakOfFrogs string) int {
	cnt := [len(previous)]int{}
	for _, ch := range croakOfFrogs {
		pre := previous[ch] // pre 为 ch 在 "croak" 中的上一个字母
		if cnt[pre] > 0 {   // 如果有青蛙发出了 pre 的声音
			cnt[pre]-- // 复用一只
		} else if ch != 'c' { // 否则青蛙必须从 'c' 开始蛙鸣
			return -1 // 不符合要求
		}
		cnt[ch]++ // 发出了 ch 的声音
	}
	if cnt['c'] > 0 || cnt['r'] > 0 || cnt['o'] > 0 || cnt['a'] > 0 {
		return -1 // 有发出其它声音的青蛙，不符合要求
	}
	return cnt['k'] // 最后青蛙们都发出了 'k' 的声音
}

func TestFrogs(t *testing.T) {
	fmt.Println(len(previous))
	//var croakOfFrogs = "croakcroakcroac"
	var croakOfFrogs = "crcoakroakcrcoakroak"
	fmt.Println(minNumberOfFrogs(croakOfFrogs))
	fmt.Println(minNumberOfFrogs2(croakOfFrogs))

}
