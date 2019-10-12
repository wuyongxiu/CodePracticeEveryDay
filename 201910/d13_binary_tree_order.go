package _01910

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(root *TreeNode) []int {
	v := make([]int, 0)
	v = append(v, root.Val)
	if root.Left != nil {
		leftV := PreOrder(root.Left)
		v = append(v, leftV...)
	}
	if root.Right != nil {
		rightV := PreOrder(root.Right)
		v = append(v, rightV...)
	}
	return v
}

func InOrder(root *TreeNode) []int {
	v := make([]int, 0)
	if root.Left != nil {
		leftV := InOrder(root.Left)
		v = append(v, leftV...)
	}
	v = append(v, root.Val)
	if root.Right != nil {
		rightV := InOrder(root.Right)
		v = append(v, rightV...)
	}
	return v
}

func PostOrder(root *TreeNode) []int {
	v := make([]int, 0)
	if root.Left != nil {
		leftV := PostOrder(root.Left)
		v = append(v, leftV...)
	}
	if root.Right != nil {
		rightV := PostOrder(root.Right)
		v = append(v, rightV...)
	}
	v = append(v, root.Val)
	return v
}
