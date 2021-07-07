// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
package main

import "fmt"

func cuttingRope(n int) int {

	l := make([]int, n)

	for i, _ := range l {
		curMax := 1
		for j:=0; j<(i+1)/2; j++ {
			curRes := (j+1) * l[i-j-1]
			if curRes < (j+1) * (i-j) {
				curRes = (j+1) * (i-j)
			}
			if curMax < curRes {
				curMax = curRes
			}
		}
		l[i] = curMax
	}
	return l[n-1]
}

func cuttingRopeII(n int) int {
	if n < 3 {
		return 1
	}
	if n == 3 {
		return 2
	}
	res := 1
	for n > 4 {
		res *= 3
		res %= 1000000007
		n -= 3
	}
	return (n * res) % 1000000007
}


func main() {
	fmt.Println(cuttingRopeII(8))
}