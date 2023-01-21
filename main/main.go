package main

import (
	"fmt"
	"lemin"
	"os"
	"strings"
)

var Route []lemin.Vertice
var filename string = os.Args[1]

func main() {
	data, _ := os.ReadFile(filename)
	content := strings.Split(string(data), "\n")
	MyMap, index := lemin.ValidateRooms(content)
	MyMap = lemin.ValidateLinks(content[index:], &MyMap)

	for k := range MyMap {
		if k.Start {
			lemin.RecursivePathFinder(k, Route)
		}
	}

	// fmt.Println(len(lemin.AllPaths))
	// for _, v := range lemin.AllPaths {
	// 	for i, item := range v {
	// 		fmt.Printf("Turn %d: %s\n", i+1, item.Name)
	// 	}
	// 	fmt.Println()
	// 	fmt.Println()
	// }
	Result := lemin.ChoosePath(lemin.CombinePaths(lemin.AllPaths))

	// for i, v := range Result {
	// 	fmt.Printf("MaxFlow %d: ", i+1)
	// 	for I, element := range v {
	// 		fmt.Println("\nPath ", I+1, ":")
	// 		for index, item := range element {
	// 			fmt.Printf("Turn %d: %s\t", index+1, item.Name)
	// 		}
	// 		fmt.Println()
	// 	}
	// 	fmt.Println()
	// 	fmt.Println()
	// }

	for i, v := range Result {
		fmt.Println("\nPath ", i+1, ":")
		for index, item := range v {
			fmt.Printf("Node %d: %s  ", index+1, item.Name)
		}
		fmt.Println()
	}

	// for k := range MyMap {
	// 	fmt.Println(k.Name)
	// 	for _, c := range k.Links {
	// 		fmt.Println("Links: ", c.Name)
	// 	}
	// 	fmt.Println()
	// }
}
