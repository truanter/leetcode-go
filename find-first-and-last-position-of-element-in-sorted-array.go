package main

import "fmt"

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	leftIdx := binarySearch(nums, target, true)
	rightIdx := binarySearch(nums, target, false) - 1
	if leftIdx <= rightIdx && rightIdx < len(nums) && nums[leftIdx] == target && nums[rightIdx] == target {
		return []int{leftIdx, rightIdx}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, lower bool) int {
	left, right, res := 0, len(nums) - 1, len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target || lower && nums[mid] >= target {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}
	return res
}

func TestFindFirstAndLast() {
	res := searchRange([]int{5,7,7,8,8,10}, 8)
	fmt.Printf("%v", res)
}