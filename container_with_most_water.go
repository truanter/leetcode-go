package main

import "fmt"
/*
https://leetcode-cn.com/problems/container-with-most-water/submissions/
*/

func maxArea(height []int) int {
	beginIndex, endIndex, res := 0, len(height)-1, 0
	for beginIndex <= endIndex {
		begin, end := height[beginIndex], height[endIndex]
		curRes := 0
		if begin > end {
			curRes = end * (endIndex - beginIndex)
			endIndex--
		} else {
			curRes = begin * (endIndex - beginIndex)
			beginIndex++
		}
		if res < curRes {
			res = curRes
		}
	}
	return res
}

func TestMaxArea() {
	fmt.Print(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
