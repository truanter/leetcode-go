package main

/**
* Definition for a binary tree node.
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}
	mediaIdx := 0
	for ; mediaIdx < len(inorder); mediaIdx ++ {
		if inorder[mediaIdx] == rootVal {
			break
		}
	}
	root.Left = buildTree(preorder[1: mediaIdx+1], inorder[:mediaIdx])
	root.Right = buildTree(preorder[mediaIdx+1:], inorder[mediaIdx+1:])
	return root
}
