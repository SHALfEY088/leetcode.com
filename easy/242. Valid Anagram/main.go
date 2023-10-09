// https://leetcode.com/problems/valid-anagram/
package main

import "fmt"

func main() {
	s, t := "anagram", "nagaram"
	s2, t2 := "rat", "car"

	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram(s2, t2))
}

//func isAnagram(s string, t string) bool {
//	sRune := []rune(s)
//	tRune := []rune(t)
//
//	if len(sRune) != len(tRune) {
//		return false
//	}
//
//	ms := make(map[rune]int)
//	mt := make(map[rune]int)
//
//	for i := 0; i < len(sRune); i++ {
//		ms[sRune[i]]++
//		mt[tRune[i]]++
//	}
//
//	for k, v := range ms {
//		if valt, exist := mt[k]; !exist || valt != v {
//			return false
//		}
//
//	}
//
//	return true
//}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}
