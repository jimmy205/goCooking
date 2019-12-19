package main

import (
	"log"
	"strings"
)

func main() {

	strings.ToLower("AAA")

	// answer = [0,1]
	towSumAn := towSum([]int{2, 5, 11, 15}, 7)
	log.Println("towSumAn : ", towSumAn)

	// answer = 3
	jewelAn := numJewelsInStones("aA", "aAAbbbb")
	log.Println("jewelAn : ", jewelAn)
}

func towSum(nums []int, target int) []int {

	// map[nums的值]nums的index
	resMap := make(map[int]int)

	/*
	 * 利用map的特性把有計算過的餘數記下來比對
	 */
	for i := range nums {
		// 比對之前算過的餘數是不是有存在MAP
		value, exists := resMap[target-nums[i]]
		// 如果有，代表現在的數字加之前的餘數 = 目標數
		if exists {
			return []int{value, i}
		}
		// 如果沒有，把現在的數字記下來，值設定為index
		resMap[nums[i]] = i
	}
	return []int{0, 0}
}

func numJewelsInStones(J string, S string) int {
	var count int
	jewelMap := make(map[byte]bool)
	for i := range J {
		jewelMap[J[i]] = true
	}

	for i := range S {
		_, ok := jewelMap[S[i]]
		if ok {
			count++
		}
	}

	// 解法2
	// var count int

	// for i := range J {
	//     n := strings.Count(S,string(J[i]))
	//     count += n
	// }

	return count
}
