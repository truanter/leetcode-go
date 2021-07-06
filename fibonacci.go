// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
package main

import "fmt"

func fib(n int) int {
	l := []int{0, 1}
	for len(l) <= n {
		l = append(l, (l[len(l)-1] + l[len(l)-2]) % 1000000007)
	}
	return l[n] % 1000000007
}

func main() {
	fmt.Println(fib(95))
}
