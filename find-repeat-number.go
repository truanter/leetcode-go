package main

import "fmt"

func findRepeatNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		if nums[abs(v)] < 0 {
			return abs(v)
		} else {
			nums[abs(v)] *= -1
		}
	}
	return res
}

func abs(v int) int {
	if v < 0 {
		return 0 - v
	}
	return v
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}