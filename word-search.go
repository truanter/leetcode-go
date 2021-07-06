package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for x:=0; x<len(board); x++ {
		for y:=0; y<len(board[0]); y++ {
			if dfs(board, x, y, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, x, y, k int, word string) bool {
	if k >= len(word) {
		return true
	}

	if x < 0 || y < 0 || x >=  len(board) || y >= len(board[0]) {
		return false
	}

	if board[x][y] == byte(0) || board[x][y] != word[k] {
		return false
	}
	tmp := board[x][y]

	board[x][y] = byte(0)

	if dfs(board, x+1, y, k+1, word) || dfs(board, x-1, y, k+1, word) || dfs(board, x, y+1, k+1, word) || dfs(board, x, y-1, k+1, word) {
		return true
	}

	board[x][y] = tmp
	return false
}

func main() {

	b := [][]byte {
		[]byte{byte('A'), byte('B'), byte('C'), byte('E')},
		[]byte{byte('S'), byte('F'), byte('C'), byte('S')},
		[]byte{byte('A'), byte('D'), byte('E'), byte('F')},
	}
	fmt.Println(exist(b, "ABCCED"))
}