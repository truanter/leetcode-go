//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
package main

func removeNthFromEnd(head *IntListNode, n int) *IntListNode {
	var slow *IntListNode
	fast := head
	distance := 1
	for fast.Next != nil {
		fast = fast.Next
		if distance < n {
			distance++
		} else {
			if slow == nil {
				slow = head
			} else {
				slow = slow.Next
			}
		}
	}
	if slow == nil {
		head = head.Next
	} else {
		if slow.Next != nil{
			slow.Next = slow.Next.Next
		}
	}
	return head
}


func TestRemoveNthFromEnd() {
	l := []int{1,2}
	n := 2
	h := GenerateIntListNode(l)
	PrintIntListNode(h)
	h = removeNthFromEnd(h, n)
	PrintIntListNode(h)
}