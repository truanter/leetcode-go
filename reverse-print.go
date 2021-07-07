package main

import "fmt"

type intListNode struct {
	Val  int
	Next *intListNode
}

func generateIntListNode(l []int) *intListNode {
	head := &intListNode{}
	cur := head
	if len(l) > 0 {
		head.Val = l[0]
	}
	for _, v := range l[1:] {
		cur.Next = &intListNode{
			Val: v,
		}
		cur = cur.Next
	}
	return head
}


func reversePrint(head *intListNode) []int {
	if head == nil {
		return make([]int, 0)
	}
	nextRes := reversePrint(head.Next)
	nextRes = append(nextRes, head.Val)
	return nextRes
}

func main() {
	fmt.Println(reversePrint(generateIntListNode([]int{1, 3, 2})))
}