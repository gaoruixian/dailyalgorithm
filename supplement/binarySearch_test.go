package supplement

import (
	"fmt"
	"testing"
)

func Test(t *testing.T)  {
									// 0 1 2 3 4  5  6
	r,v := findNumInSortedArrays([]int{1,3,5,7,8,11,21},21)
	if r {
		fmt.Println(v.(int))
	}else {
		fmt.Println("不存在")
	}
}

/**
有序数组中查找整数
0  1  2  3  4
[2,3,12,19,21]
 */
func findNumInSortedArrays(nums []int,num int) (bool,interface{})  {
	if num < nums[0] || num > nums[len(nums)-1] {
		return false,nil
	}
	// begin end
	begin,end,mid := 0,len(nums)-1,0
	for begin <= end {
		mid = (begin + end) >> 1
		if nums[mid] == num {
			return true,mid
		}else if nums[mid] > num {
			end = mid - 1
		}else if nums[mid] < num{
			begin = mid + 1
		}
	}
	return false,nil
}