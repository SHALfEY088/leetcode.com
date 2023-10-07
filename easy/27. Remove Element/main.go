// https://leetcode.com/problems/remove-element/description/
package main

import "fmt"

func main() {

	n1 := []int{3, 2, 2, 3}
	fmt.Println(removeElement(n1, 3), "nums = ", n1)
	n2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(n2, 2), "nums = ", n2)

}

func removeElement(nums []int, val int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
