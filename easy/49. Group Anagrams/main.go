https://leetcode.com/problems/group-anagrams/description/
package main

import "fmt"

func main() {
	fmt.Println([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println([]string{""})
	fmt.Println([]string{"a"})
}

func groupAnagrams(strs []string) [][]string {
	return make([][]string, 0)
}
