package _01910

import (
	"testing"
)

func TestOrder(t *testing.T) {
	/**
		   0
		1    2
	  11 10

	*/

	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 11,
			},
			Right: &TreeNode{
				Val: 10,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	t.Log("PreOrder：", PreOrder(root))
	t.Log("InOrder：", InOrder(root))
	t.Log("PostOrder：", PostOrder(root))
}
