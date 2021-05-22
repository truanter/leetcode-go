package main


func xorGame(nums []int) bool {
	if len(nums) % 2 == 0 {
		return true
	}
	res := 0
	for _, v := range nums {
		res = res ^ v
	}
	return res == 0
}


