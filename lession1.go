package main

import (
	"fmt"
	"lession/tree"
)

type I interface {
	Hello()
}

type M struct {
}

func (m M) Hello() {
	fmt.Println("Hello!")
}

func main() {
	node := tree.TreeNode{Value: 0}
	node.Left = &tree.TreeNode{
		Value: 1,
		Right: &tree.TreeNode{Value: 2},
	}
	node.Right = &tree.TreeNode{
		Value: 3,
		Left:  &tree.TreeNode{Value: 4},
	}
	node.PostOrder()
}

/*
1. fib数列写入fib.txt，使用defer
2. 对file exsits进行错误处理，使用pathError
3. 服务器读取文件，并进行错误处理
*/
