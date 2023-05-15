package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.
 */

func Test(t *testing.T)  {
	nums1 := []int{3,5,9,10}
	nums2 := []int{8,16}
	fmt.Printf("result:%v\n",medianOfTwoSortedArrays(nums1,nums2))
	fmt.Printf("result:%v\n",findMedianOfTwoSortedArrays(nums1,nums2))
	fmt.Printf("result:%v\n",findMedianOfTwoSortedArraysByBinarySearch(nums1,nums2))


}
func medianOfTwoSortedArrays(nums1,nums2 []int) float32 {
	// 合并 space complexity O(m+n)
	var all = append(nums1,nums2...)
	// 排序
	//sort.Ints(all) 冒泡排序 time complexity O(n^2)
	for i:=0; i < len(all); i++  {
		for j :=0; j < len(all)-1-i ; j++  {
			if all[j] > all[j+1] {
				temp := all[j+1]
				all[j+1] = all[j]
				all[j] = temp
			}
		}
	}
	fmt.Println(all)
	fmt.Println(len(all) & 1)
	// even
	if (len(all) & 1) == 0 {
		return float32(all[len(all) >> 1 - 1] + all[len(all) >> 1])/2
	}else {
		return float32(all[(len(all))/2])
	}
/**
4,7,8,2,1
 */
}
func findMedianOfTwoSortedArrays(nums1,nums2 []int) float32  {
	// space complexity O(m+n)
	var all = make([]int,len(nums1) + len(nums2))
	var(
		i = 0
		j = 0
		index = 0
	)
	// time complexity O(m+n)
	for (i < len(nums1)) && (j < len(nums2)) {
		if nums1[i] < nums2[j]{
			all[index] = nums1[i]
			i++
			index++
		}else if nums1[i] > nums2[j] {
			all[index] = nums2[j]
			j++
			index++
		}else {
			all[index] = nums1[i]
			i++
			index++
		}
	}
	for i < len(nums1)  {
		all[index] = nums1[i]
		i++
	}
	for j < len(nums2) {
		all[index] = nums2[j]
		j++
	}
	if len(all) & 1 ==0 {
		return float32(all[len(all) >> 1 - 1] + all[len(all) >> 1])/2
	}else {
		return float32(all[(len(all) + 1)/2])
	}
}
/**
findMedianOfTwoSortedArrays方法的优化，将space complexity由O(m+n) --> O(1)
通过指针虚拟出来m+n个空间
 */
func findMedianOfTwoSortedArrays2(nums1,nums2 []int) float32 {
	/*
		nums1 := []int{3,5,9,10,12}
		nums2 := []int{8,16}
	 */
	p1,p2,midPIndex := 0,0,(len(nums1) + len(nums2)) >> 1
	p3 := p1
	oddNumber := ((len(nums1) + len(nums2)) & 1) == 1
	even1,even2 := 0,0
	/**
	nums1 := []int{3,5,9,10,12}
	nums2 := []int{8,16,17}

	 */
	for p3 < midPIndex {
		fmt.Printf("循环:%v\n",p3)
		if nums1[p1] < nums2[p2] {
			p1++
			p3++
			if p3 == midPIndex && oddNumber {
				return float32(nums1[p1])
			}else if !oddNumber && (p3 == midPIndex-1) {
				even1 = nums1[p1]
			}else if !oddNumber && p3 == midPIndex {
				even2 = nums1[p1]
			}
		}else {
			p2++
			p3++
			if p3 == midPIndex && oddNumber {
				return float32(nums2[p2])
			}else if !oddNumber && (p3 == midPIndex-1) {
				even1 = nums2[p2]
			}else if !oddNumber && p3 == midPIndex {
				even2 = nums2[p2]
			}
		}
	}
	return float32(even1 + even2) / 2
}
/**
findMedianOfTwoSortedArrays方法的优化，将space complexity由O(m+n) --> O(1)
由于数组a,b确定，合并后的数组中位数位置是确定的。
m+n为奇数时，
*/
func findMedianOfTwoSortedArrays22(nums1,nums2 []int) float32 {
	/*
		nums1 := []int{3,5,9,10}
		nums2 := []int{8,16}
	*/
	p1,p2,midPIndex := 0,0,(len(nums1) + len(nums2)) >> 1
	left,right := 0,0
	oddNumber := ((len(nums1) + len(nums2)) & 1) == 1
	/**
	1,2,3,4,5,6,7,8
	1,2,3,4,5,6,7,8,9
	nums1 := []int{3,5,9,10}
	nums2 := []int{8,16,17}
	3,4,8,9,10,16,17
 	1 2 3 4 5 6 7
	i =0 --> left = 0     , right = 3, p1 = 1,p2=0;
	i =1 --> left = 3     , right = 5, p1 = 2,p2=0;
	i =2 --> left = 5     , right = 8, p1 = 2,p2=1
	i =3 --> left = 8     , right = 9, p1 = 3,p2=1
	i =4 --> left = 9     , right = 10,p1 = 4,p2=1
	*/
	for i:=0;i < midPIndex + 1;i++{
		left = right
		if p1 < len(nums1) && (nums1[p1] < nums2[p2] || p2 >= len(nums2)) {
			right = nums1[p1]
			p1++
		}else {
			right = nums2[p2]
			p2++
		}
	}
	if oddNumber {
		return float32(right)
	}else {
		return float32(left + right) / 2
	}
}

func findMedianOfTwoSortedArrays3(nums1,nums2 []int) float32  {
	if len(nums1) > len(nums2) {
		findMedianOfTwoSortedArrays3(nums2,nums1)
	}
	/**
	nums1 := []int{3,5,9,10,12}
	nums2 := []int{8,16,17}
	 */
	p1,p2,allSize := 0,0,len(nums1) + len(nums2)
	oddNumber := (allSize & 1) == 1
	left,right := -1,-1
	for i:=0;i<=(allSize >> 1);i++ {
		left = right
		if p1 < len(nums1) && ((nums1[p1] < nums2[p2]) || p2 >= len(nums2)) {
			right = nums1[p1]
			p1++
		}else {
			right = nums2[p2]
			p2++
		}
	}
	if oddNumber {
		return float32(right)
	}else {
		fmt.Printf("left:%v",left)
		fmt.Printf("left:%v",right)

		return float32(left + right) / 2
	}

}

func max(a,b int) int  {
	if a > b {
		return a
	}
	return b
}

func findMedianOfTwoSortedArraysByBinarySearch(nums1,nums2 []int) float64  {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianOfTwoSortedArraysByBinarySearch(nums2, nums1)
	}
	/**
	nums1 := []int{3,5,9,10,12}
	nums2 := []int{8,16,17}
	 */
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
			//nums1Mid = high + 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
			//nums1Mid = low - 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = int(math.Max(float64(nums1[nums1Mid-1]), float64(nums2[nums2Mid-1])))
	}
	// 奇数
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = int(math.Min(float64(nums1[nums1Mid]), float64(nums2[nums2Mid])))
	}
	return float64(midLeft+midRight) / 2
}
func TestY(t *testing.T)  {
	fmt.Println(1 + (9-1)>>2)
}
func TestYiWei(t *testing.T)  {
	//var a = 4
	//fmt.Println(a >> 2)
	//fmt.Println(a >> 2 + 1)
	//fmt.Println(float32(17 / 2))
	//
	//fmt.Println(1 + 15 >> 4)
	//
	//fmt.Println(6 >>1 )

	nums1 := []int{3,5,9,10}
	nums2 := []int{8}

	//fmt.Printf("result:%v\n",findMedianOfTwoSortedArrays2(nums1,nums2))
	fmt.Printf("result:%v\n",medianOfTwoSortedArrays(nums1,nums2))

	fmt.Printf("result:%v\n",findMedianOfTwoSortedArrays22(nums1,nums2))
	fmt.Printf("result:%v\n",findMedianOfTwoSortedArrays3(nums1,nums2))

	//fmt.Printf("result:%v\n",findMedianOfTwoSortedArrays3(nums1,nums2))
	//fmt.Printf("result:%v\n",medianOfTwoSortedArrays(nums1,nums2))

}