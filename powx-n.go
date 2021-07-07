package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	isMinus := n < 0
	isOdds := n % 2 != 0
	k := n
	if isMinus {
		k = 0-n
	}
	res := myPow(x*x, k/2)
	if isOdds {
		res *= x
	}
	if isMinus {
		res = 1/res
	}
	return res
}

func main() {
	fmt.Println(myPow(2.0, -2))
}