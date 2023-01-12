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
	//result := lemin.LinkedList(MyMap)

	// fmt.Println("head: ", result.Head, "\t tail: ", result.Tail)
	for k := range MyMap {
		fmt.Print("key:", k)
		for i, v := range k.Neighbour {
			fmt.Println("neighbour num %d is: %v", i, *v)
		}
		fmt.Println()
	}
}
