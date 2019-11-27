package main

import (
	"github.com/chez-shanpu/repo2tree/model"
	"io/ioutil"
	"log"
	"path/filepath"
)

func MakeLayer(dirPaths []string, depth int, parentNode *model.Node) *model.Node {
	var rightmostNode *model.Node
	var leftmostNode *model.Node
	var targetNode *model.Node

	for _, dirPath := range dirPaths {
		targetNode = makeNode(dirPath, depth, parentNode)
		if leftmostNode == nil {
			leftmostNode = targetNode
		} else {
			rightmostNode.NextNode = targetNode
		}
		rightmostNode = targetNode
	}
	return leftmostNode
}

func makeNode(dirPath string, depth int, parentNode *model.Node) *model.Node {
	var node model.Node
	var nodeDataIndex int
	var subDirPaths []string

	_, node.DirectoryName = filepath.Split(dirPath)
	node.Data = [9]float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			subDirPaths = append(subDirPaths, filepath.Join(dirPath, file.Name()))
		} else {
			nodeDataIndex = FileClassifier(file.Name())
			if node.Data[nodeDataIndex] == 0 {
				node.Data[nodeDataIndex] = 1
			}
		}
	}
	for key := range node.Data {
		if parentNode == nil {
			node.Data[key] = node.Data[key] / float64(depth)
		} else {
			node.Data[key] = node.Data[key]/float64(depth) + parentNode.Data[key]
		}
	}
	if subDirPaths != nil {
		node.ChildNode = MakeLayer(subDirPaths, depth+1, &node)
	}
	return &node
}
