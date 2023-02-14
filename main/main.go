package main

import (
	"fmt"
	"lemin"
	"os"
	"strings"
)

var Route []lemin.Vertice
var filename string = os.Args[1]
var NumberAnts int

func main() {
	//data, _ := os.ReadFile(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
	  log.Errorf(err)
	content := strings.Split(string(data), "\n")

	NumberAnts = lemin.ValidateAnts(content)

	MyMap, index := lemin.ValidateRooms(content)

	MyMap = lemin.ValidateLinks(content[index:], &MyMap)

	for k := range MyMap {
		if k.Start {
			lemin.RecursivePathFinder(k, Route)
		}
	}

	Result := lemin.ChoosePath(lemin.CombinePaths(lemin.AllPaths))
	lemin.QueueThem(NumberAnts, Result)

	for i, v := range Result {
		fmt.Println("\nPath ", i+1, ":")
		for index, item := range v {
			fmt.Printf("Node %d: %s  ", index+1, item.Name)
		}
		fmt.Println()
	}
	// for i, v := range lemin.AllPaths {
	// 	fmt.Println("Path ", i+1, ":")
	// 	for index, element := range v {
	// 		fmt.Println("Node ", index+1, " is ", element.Name)
	// 	}
	// 	fmt.Println()

	// }
}
