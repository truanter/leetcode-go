//https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
package main

import "fmt"

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	cur, pre := 2, 1
	for i:=3; i<=n; i++ {
		pre, cur = cur, (cur + pre) % 1000000007
	}
	return cur
}

func main() {
	fmt.Println(numWays(7))
}