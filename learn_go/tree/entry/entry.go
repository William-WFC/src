package main

import (
	"fmt"
	"learnGo/tree"
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

func add(x int) {

}

func init() {}
func init() {}
func init() {}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(n *tree.Node) {
		nodeCount++
	})
	fmt.Println("node count:", nodeCount)

	x := []int{1, 2, 3, 4}
	y := x[:2]
	y = append(y, 5)
	fmt.Println(x)

	// myRoot := myTreeNode{&root}
	// myRoot.postOrder()
	ages := make(map[int][]string)
	fmt.Println(ages)

}
