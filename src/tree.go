package main

import (
	"github.com/chez-shanpu/repo2tree/src/model"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"
)

func MakeLayer(dirPaths []string, depth int, parentNode *model.Node) *model.Node {
	var leftmostNode *model.Node
	var targetNode *model.Node

	for _, dirPath := range dirPaths {
		targetNode = makeNode(dirPath, depth, parentNode)
		if leftmostNode == nil {
			leftmostNode = targetNode
		} else {
			leftmostNode = insertNode(leftmostNode, targetNode)
		}
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

func insertNode(leftmostNode *model.Node, targetNode *model.Node) *model.Node {
	var targetNodeValue float64
	var nextNodeValue float64
	var leftmostNodeValue float64
	var node *model.Node

	for key, val := range targetNode.Data {
		targetNodeValue += math.Pow(val, math.Pow(10, float64(key)))
	}
	for key, val := range leftmostNode.Data {
		leftmostNodeValue += math.Pow(val, math.Pow(10, float64(key)))
	}

	if targetNodeValue < leftmostNodeValue {
		targetNode.NextNode = leftmostNode
		leftmostNode = targetNode
		return leftmostNode
	}
	for node = leftmostNode; node.NextNode != nil; node = node.NextNode {
		for key, val := range node.NextNode.Data {
			nextNodeValue += math.Pow(val, math.Pow(10, float64(key)))
		}
		if targetNodeValue < nextNodeValue {
			targetNode.NextNode = node.NextNode
			node.NextNode = targetNode
			return leftmostNode
		}
	}
	node.NextNode = targetNode
	return leftmostNode
}
