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


func main() {
	for i:=1; i< 10; i++ {
		fmt.Println(fmt.Sprintf("n=%d, res=%d", i, cuttingRope(i)))
	}
}