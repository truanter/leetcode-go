//https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/

package main

type MinStack struct {
	stack []int
	minIdx int
	lastMin []int
}


/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		stack: make([]int, 0),
		minIdx: -1,
		lastMin: make([]int, 0),
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	this.lastMin = append(this.lastMin, this.minIdx)
	if this.minIdx >= 0 {
		if x < this.stack[this.minIdx] {
			this.minIdx = len(this.stack) -1
		}
	} else {
		this.minIdx = 0
	}
}


func (this *MinStack) Pop()  {
	this.minIdx = this.lastMin[len(this.lastMin)-1]
	this.lastMin = this.lastMin[:len(this.lastMin)-1]
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) Min() int {
	return this.stack[this.minIdx]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */