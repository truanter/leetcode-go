//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

package main

type CQueue struct {
	A []int
	B []int
}


func Constructor() CQueue {
	return CQueue{
		A: make([]int, 0),
		B: make([]int, 0),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.A = append(this.A, value)
}


func (this *CQueue) DeleteHead() int {
	var v int
	if len(this.B) > 0 {
		v, this.B = this.B[0], this.B[1:]
		return v
	} else {
		if len(this.A) < 1 {
			return -1
		}
		for len(this.A) > 0 {
			v, this.A = this.A[0], this.A[1:]
			this.B = append(this.B, v)
		}
		v, this.B = this.B[0], this.B[1:]
		return v
	}
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
