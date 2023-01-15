package main

import (
	"fmt"
	"lemin"
	"os"
	"strings"
)

var filename string = os.Args[1]

func main() {
	data, _ := os.ReadFile(filename)
	content := strings.Split(string(data), "\n")
	MyMap, index := lemin.ValidateRooms(content)
	MyMap = lemin.ValidateLinks(content[index:], &MyMap)
	lemin.AdjustMap(&MyMap)

	for k := range MyMap {
		fmt.Println(k.Name)
		for _, p := range k.Parents {
			fmt.Println("Parents: ", p.Name)
		}
		for _, c := range k.Children {
			fmt.Println("Children: ", c.Name)
		}
		fmt.Println()
	}
	// for k := range MyMap {
	// 	fmt.Printf("Room: %s\n", k.Name)
	// 	for _, v := range k.Children {
	// 		fmt.Println("Relative neighbour(s): ", *v)
	// 	}
	// 	fmt.Println()
	// }
}
