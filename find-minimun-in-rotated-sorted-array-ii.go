// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/submissions/
package main


func minArray(numbers []int) int {
	res := numbers[0]

	for index, v := range numbers[1:] {
		if v < numbers[index] {
			return v
		}
	}
	return res
}