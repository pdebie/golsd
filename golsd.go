package main

import (
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net/http"
)

// Node represents a node in a directory tree. FullPath and Info are available but not encoded into JSON
type Node struct {
	Filename string       `json:"filename"`              //File or Directory name serialized as title
	Size     int64        `json:"size"`                //Size
	Folder   bool         `json:"folder"`   //Is this a folder - ie. Info.IsDir()
}

// var myNodes = map[string]*Node{}
var nodes []*Node

func readFiles() {

    files, _ := ioutil.ReadDir("./")
    for _, f := range files {
        m := &Node{Filename: f.Name(), Size: f.Size(), Folder: f.IsDir()}
        nodes = append(nodes, m)
    }
}

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        readFiles()
        c.JSON(http.StatusOK, nodes) 
    })
    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}