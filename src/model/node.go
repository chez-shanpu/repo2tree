package model

type Node struct {
	Data      [9]float64 `json:"data"`
	NextNode  *Node      `json:"next_node"`
	ChildNode *Node      `json:"child_node"`
}

type NodeInfo struct {
	RootNode       *Node  `json:"root_node"`
	RepositoryName string `json:"repository_name"`
	Language       string `json:"language"`
	CreatedDate        string `json:"created_date"`
}
