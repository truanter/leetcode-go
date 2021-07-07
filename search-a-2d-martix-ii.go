package main

import "fmt"

func findNumberIn2DArrayII(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	if len(matrix[0]) == 0 {
		return false
	}

	x, y := 0, len(matrix[0]) -1

	for y>=0 && x < len(matrix) {
		t := matrix[x][y]
		if t == target {
			return true
		}
		if t > target {
			y -= 1
		} else {
			x += 1
		}
	}
	return false
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	if len(matrix[0]) == 0 {
		return false
	}
	return find(0, 0, len(matrix[0]), len(matrix), target, matrix)
}

func find(x, y, w, h, target int, matrix [][]int) bool {
	if x >= w || y >= h {
		return false
	}

	if x == w-1 && y == h-1 {
		return matrix[y][x] == target
	}

	for curX:=x; curX<w; curX++{
		if matrix[y][curX] == target {
			return true
		}
		if matrix[y][curX] > target {
			w = curX
			break
		}
	}

	for curY:=y; curY<h; curY++ {
		if matrix[curY][x] == target {
			return true
		}
		if matrix[curY][x] > target {
			h = curY
			break
		}
	}
	return find(x+1, y+1, w, h, target, matrix)
}

func main() {
	m2 := [][]int {
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	_ = [][]int {
		{1, 2},
	}

	fmt.Println(findNumberIn2DArrayII(m2, 20))
}