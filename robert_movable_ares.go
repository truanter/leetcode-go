// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
package main

import "fmt"

func movingCount(m int, n int, k int) int {
	table := make([][]*int, m)
	for i := range table {
		table[i] = make([]*int, n)
	}
	res := count(table, m, n, 0, 0, k)
	return res
}

func count(table [][]*int, m, n, x, y, k int) (res int) {
	if x >= m || y >= n || x < 0 || y < 0{
		return
	}

	if table[x][y] != nil {
		return
	}

	if getBitSum(x) + getBitSum(y) <= k {
		res += 1
	} else {
		return
	}

	j := 1

	table[x][y] = &j

	res += count(table, m, n, x+1, y, k) +
		count(table, m, n, x, y+1, k)
	return res
}

func getBitSum(num int) int {
	res := num % 10
	b := num / 10
	for b>0 {
		res += b % 10
		b = b / 10
	}
	return res
}

func main() {
	fmt.Println(movingCount(16, 8, 4))
}