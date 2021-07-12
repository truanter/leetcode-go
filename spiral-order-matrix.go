package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	if len(matrix[0]) == 0 {
		return res
	}

	i, j, m, n := 1, 0, len(matrix), len(matrix[0])
	x, y := 0, 0
	xDelta, yDelta := 0, 1

	for (i<m || j<n ) && len(res) < len(matrix)*len(matrix[0]) {
		res = append(res, matrix[x][y])

		if y == n-1 && yDelta > 0{
			xDelta = 1
			yDelta = 0
			n--
		}

		if y == j && yDelta < 0{
			xDelta = -1
			yDelta = 0
			j++
		}

		if x == m-1 && xDelta > 0{
			yDelta = -1
			xDelta = 0
			m--
		}

		if x == i && xDelta<0{
			yDelta = 1
			xDelta = 0
			i++
		}

		x += xDelta
		y += yDelta

	}
	return res
}

func main() {
	f := spiralOrder([][]int {
		{1, 2, 3, 0, 1},
		{4, 5, 6, 0, 2},
		{7, 8, 9, 0, 3},
		{0, 0, 0, 0, 4},
	})
	fmt.Println(f)
}