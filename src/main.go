package main

import (
	"encoding/json"
	"github.com/chez-shanpu/repo2tree/src/model"
	"log"
	"os"
	"time"
)

func main() {
	var repoRootPath string
	var repositoryName string
	var language string
	var createdDate string
	var outputName string
	var rootNode *model.Node

	createdDate = time.Now().String()
	treeRoot := []string{repoRootPath}
	rootNode = MakeLayer(treeRoot, 1, nil)
	nodeInfo := model.NodeInfo{
		RootNode:       rootNode,
		RepositoryName: repositoryName,
		Language:       language,
		CreatedDate:    createdDate,
	}
	outputJson(outputName, nodeInfo)
	log.Printf("Output to %v completed", outputName)
}

func outputJson(outputName string, nodeInfo model.NodeInfo) {
	file, err := os.Create(outputName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, _ := json.Marshal(nodeInfo)
	file.Write(bytes)
}
