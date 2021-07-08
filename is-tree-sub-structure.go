// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}

	if A == nil {
		return false
	}

	return isSameRootBigger(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)

}

func isSameRootBigger(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isSameRootBigger(a.Left, b.Left) && isSameRootBigger(a.Right, b.Right)
}
