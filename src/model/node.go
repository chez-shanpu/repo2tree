package model

type Node struct {
	Data      [9]float64
	NextNode  *Node
	ChildNode *Node
}
