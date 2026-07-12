package tree

import (
	"fmt"
)

func (treeNode *TreeNode) PreOrder() {

	if treeNode == nil {
		return
	}

	fmt.Printf("Node Val: %d\n", treeNode.Val)
	treeNode.Left.PreOrder()
	treeNode.Right.PreOrder()
}
