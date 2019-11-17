package main

import (
	"encoding/json"
	"flag"
	"github.com/chez-shanpu/repo2tree/src/model"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var repoRootPath string
	var repositoryName string
	var language string
	var createdDate string
	var outputName string
	var rootNode *model.Node

	flag.StringVar(&repoRootPath, "p", "", "path to repository")
	flag.StringVar(&language, "l", "", "repository's programming language")
	flag.StringVar(&outputName, "o", "", "output file name")

	createdDate = time.Now().String()
	treeRoot := []string{repoRootPath}
	rootNode = MakeLayer(treeRoot, 1, nil)
	nodeInfo := model.NodeInfo{
		RootNode:       rootNode,
		RepositoryName: repositoryName,
		Language:       language,
		CreatedDate:    createdDate,
	}
	_, outputName = filepath.Split(repoRootPath)
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
