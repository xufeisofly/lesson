package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node *TreeNode) Print() {
	fmt.Printf("%d", node.Value)
}

func (node *TreeNode) PostOrder() {
	node.PostOrderFunc(func(node *TreeNode) {
		node.Print()
	})
}

func (node *TreeNode) PostOrderFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.PostOrderFunc(f)
	f(node)
	node.Right.PostOrderFunc(f)
}
