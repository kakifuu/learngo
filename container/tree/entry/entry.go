package main

import (
	"fmt"
	"learngo/container/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrderTraverse() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrderTraverse()
	right := myTreeNode{myNode.node.Right}
	right.postOrderTraverse()
	myNode.node.PrintValue()
}

func main() {
	var pRoot *tree.Node
	pRoot = tree.CreateNode(3)
	pRoot.Left = &tree.Node{}
	pRoot.Left.Right = tree.CreateNode(2)
	pRoot.Right = &tree.Node{Value: 5}
	pRoot.Right.Left = new(tree.Node)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, pRoot},
	}
	fmt.Println(nodes)
	pRoot.PrintValue()

	pRoot.Right.Left.PrintValue()
	pRoot.Right.Left.SetValue(4)
	pRoot.Right.Left.PrintValue()

	fmt.Println()

	nodeCount := 0
	pRoot.TraverseFunc(func(n *tree.Node) {
		nodeCount++
	})
	fmt.Println("node count:", nodeCount)
	fmt.Println()

	node := myTreeNode{node: pRoot}
	node.postOrderTraverse()

	fmt.Println()
	fmt.Println("---Traverse with channel---")
	c := pRoot.TraverseWithChannel()
	maxNodeValue := 0
	for node := range c {
		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}
	}
	fmt.Printf("Max node value: %v", maxNodeValue)
}
