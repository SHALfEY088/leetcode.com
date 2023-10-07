// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 1, 2}
	//  indexes    0  1  2  3  4  5  6  7  8  9       len(nums)==10
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} // - slice (dynamic array)
	//                            l
	//                                          r
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(nums1)
	fmt.Println(removeDuplicates(nums2))
	fmt.Println(nums2)
}

func removeDuplicates(nums []int) int {
	l := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
