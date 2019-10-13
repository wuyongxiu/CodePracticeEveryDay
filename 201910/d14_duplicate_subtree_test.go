package _01910

import "testing"



func TestFindDuplicateSubtrees(t *testing.T) {
	/**
		     1
	       / \
	      2   3
	     /   / \
	    4   2   4
	       /
	      4
	*/

	testNode := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	t.Log("fint duplicate subtrees")
	for _, duplicateNode := range FindDuplicateSubtrees(testNode) {
		t.Log(PreOrderStr(duplicateNode))
	}
	/**
	  2,4  4
	*/
}

func TestPreOrderStr(t *testing.T) {
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
	t.Log(PreOrderStr(root))
	// should be 0,1,11,10,2
}


func TestAddSubtrees(t *testing.T) {
	testNode := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	/**
		 1
	   / \
	  2   3
	 /   / \
	4   2   4
	   /
	  4
	*/

	subTrees = make(map[*TreeNode]string, 0)
	AddSubtrees(testNode)
	for _, tree := range subTrees {
		t.Log(tree)
	}

}
