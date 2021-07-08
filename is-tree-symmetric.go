package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isTwoTreeSymmetric(root.Left, root.Right)

}

func isTwoTreeSymmetric(a *TreeNode, b*TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isTwoTreeSymmetric(a.Left, b.Right) && isTwoTreeSymmetric(a.Right, b.Left)
}