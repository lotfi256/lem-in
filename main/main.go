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

	for _, v := range lemin.AllPaths {
		for i, item := range v {
			fmt.Print("  Move ", i+1, ":  ", item.Name)
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
