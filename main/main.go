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

	MyMap = lemin.ValidateLinks(content[index:], MyMap)

	fmt.Println(MyMap)
}
