package tree

import "testing"

func TestPreOrder(t *testing.T) {
	/*
	       1
	      / \
	     2   3
	    /
	   4
	*/

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	root.PreOrder()
}
