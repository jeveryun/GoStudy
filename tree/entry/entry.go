package main

import (
	"fmt"
	"study1/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 4}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Traverse()
	nodeCount := 0
	root.TraverseFunc((func(n *tree.Node) {
		nodeCount++
	}))
	fmt.Println(nodeCount)

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	fmt.Println()

	// root.Left.Right.setValue(10)
	// root.Left.Right.print()

	// root.print()
	// root.setValue(100)

	// pRoot := &root
	// pRoot.print()
	// pRoot.setValue(200)
	// pRoot.print()

	// fmt.Println()

	// var vRoot *Node
	// vRoot.setValue(200)
	// vRoot = &root
	// pRoot.setValue(300)
	// pRoot.print()

	// nodes := []Node {
	// 	{value: 3},
	// 	{},
	// 	{6,nil,&root},
	// }

	// fmt.Println(nodes)
}
