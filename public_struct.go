package main

import "fmt"

type IntListNode struct {
	Val  int
	Next *IntListNode
}

func GenerateIntListNode(l []int) *IntListNode {
	head := &IntListNode{}
	cur := head
	if len(l) > 0 {
		head.Val = l[0]
	}
	for _, v := range l[1:] {
		cur.Next = &IntListNode{
			Val: v,
		}
		cur = cur.Next
	}
	return head
}

func PrintIntListNode(node *IntListNode) {
	res := make([]int, 0)
	for node != nil && node.Next != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	if node != nil {
		res = append(res, node.Val)
	}
	fmt.Println(fmt.Sprintf("%v", res))
}
