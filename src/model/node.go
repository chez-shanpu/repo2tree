package model

type Node struct {
	Data      [10]float64
	NextNode  *Node
	ChildNode *Node
}
