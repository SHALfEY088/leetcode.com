// https://leetcode.com/problems/top-k-frequent-elements/
package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1, 2]
	//fmt.Println(topKFrequent([]int{1}, 1))                // [1]
	fmt.Println(topKFrequent([]int{1, 2}, 2)) // [1, 2]
}

func topKFrequent(nums []int, k int) (res []int) {
	cntMp := make(map[int]int)

	for _, num := range nums {
		if cnt, exist := cntMp[num]; exist {
			cntMp[num] = cnt + 1
		} else {
			cntMp[num] = 1
		}
	}

	cntSl := make([][]int, len(nums)+1)

	for num, cnt := range cntMp {
		cntSl[cnt] = append(cntSl[cnt], num)
	}

	for i := len(cntSl) - 1; i > 0; i-- {
		res = append(res, cntSl[i]...)
		if len(res) == k {
			return
		}

	}

	return
}
