package tree

import (
	"fmt"
)

func (treeNode *TreeNode) InOrder() {

	if treeNode == nil {
		return
	}

	treeNode.Left.InOrder()
	fmt.Printf("Node Val: %d\n", treeNode.Val)
	treeNode.Right.InOrder()
}
