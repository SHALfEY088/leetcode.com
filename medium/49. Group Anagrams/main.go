// https://leetcode.com/problems/group-anagrams/description/
package main

import "fmt"

func main() {
	//groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}) //["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}

func groupAnagrams(strs []string) [][]string {
	anMap := make(map[[26]int][]string)
	for _, s := range strs {

		var cnt [26]int

		for _, c := range s {
			cnt[c-'a']++
		}

		anMap[cnt] = append(anMap[cnt], s)
		//fmt.Printf("%v\n", anMap)

	}

	var res = make([][]string, len(anMap))
	idx := 0
	for _, v := range anMap {
		res[idx] = v
		idx++
	}

	//fmt.Println(res)
	return res
}
