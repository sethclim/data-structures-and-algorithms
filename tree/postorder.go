package tree

import (
	"fmt"
)

func (treeNode *TreeNode) PostOrder() {

	if treeNode == nil {
		return
	}

	treeNode.Left.PostOrder()
	treeNode.Right.PostOrder()

	fmt.Printf("Node Val: %d\n", treeNode.Val)
}
