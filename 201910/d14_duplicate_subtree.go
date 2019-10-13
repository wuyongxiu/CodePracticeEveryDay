package _01910

import (
	"strconv"
)

var subTrees map[*TreeNode]string

type DuplicateTree struct {
	Num      int
	TreeList []*TreeNode
}


//leetcode  https://leetcode-cn.com/problems/find-duplicate-subtrees/
/**
 实现思路：
 判断2个子树是否相等： 前序排序二叉树中的每个子树，并将排序后的结构以字符串的方式（逗号为分隔符）存储到slice中。通过判断字符串可以判断2个子树是否相等
 将子树的结构和出现次数存入map中（duplicateV）。map[string]num 子数序列化后的字符串 -> 出现次数
 遍历存储子数的slice, 如果该子树存在于duplicateV, 且出现次数=1， 那么将改子树加入到最终结果slice中。最后将遍历到的子树添加到 duplicateV中，且num++

 优化思路：
 1. 函数命名可读性不够强，需要优化
 2. AddSubtrees中用到的存储子树结构中，为了让FindDuplicateSubtrees可以访问到所有子树序列化后的结构，用到了全局变量，考虑是否可以用匿名变量等。


*/

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode{
	subTrees = make(map[*TreeNode]string, 0)
	duplicateTrees := make([]*TreeNode, 0)
	AddSubtrees(root)
	duplicateV := make(map[string]int)
	for node, treeOrder := range subTrees {
		num, ok := duplicateV[treeOrder]
		if ok && num == 1 {
			duplicateTrees = append(duplicateTrees, node)
		}
		duplicateV[treeOrder] = num + 1

	}
	return duplicateTrees

}

func AddSubtrees(root *TreeNode) {
	subTrees[root] = PreOrderStr(root)
	if root.Right != nil {
		AddSubtrees(root.Right)
	}
	if root.Left != nil {
		AddSubtrees(root.Left)
	}
}

func PreOrderStr(root *TreeNode) string {
	str := strconv.Itoa(root.Val)
	if root.Left != nil {
		leftV := PreOrderStr(root.Left)
		str = str + "," + leftV
	}
	if root.Right != nil {
		rightV := PreOrderStr(root.Right)
		str = str + "," + rightV
	}
	return str
}
