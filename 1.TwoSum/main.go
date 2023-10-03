package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

/*// Time Complexity: O(n^2)
// Space Complexity: O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}*/

// Time Complexity: O(n)
// Space Complexity: O(n)
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for indx, num := range nums {
		if requiredIdx, isPresent := indexMap[target-num]; isPresent {
			return []int{requiredIdx, indx}
		}
		indexMap[num] = indx
	}
	return []int{}
}
