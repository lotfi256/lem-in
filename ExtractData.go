package lemin

import (
	"log"
	"strconv"
	"strings"
)

var Start, End string = "##start", "##end"

type Map map[*Vertice]struct{}

func ValidateAnts(data []string) int {
	NumLine1, _ := strconv.Atoi(data[0])
	NumLine2, _ := strconv.Atoi(data[1])
	if NumLine2 != 0 {
		log.Fatalf("Invalid format of inputted ants at line 2")
	} else if NumLine1 <= 0 {
		log.Fatalf("Not enough ants input to validate")
	}
	return NumLine1
}

/*
Haven't handle '#' or 'L' Yet.
Initiate a counter for start and end
through a different function.
*/
func ValidateRooms(data []string) (Map, int) {
	MyMap := make(Map)
	var index int

	for i := 1; i < len(data); i++ {
		// if room detected
		if len(strings.Fields(data[i])) == 3 {

			//normal room
			_, normal := MyMap[&Vertice{Name: strings.Fields(data[i])[0]}]
			//start room
			_, start := MyMap[&Vertice{Name: strings.Fields(data[i])[0], Start: true}]
			//end room
			_, end := MyMap[&Vertice{Name: strings.Fields(data[i])[0], End: true}]
			if normal || start || end {
				log.Fatalf("Sorry, but room %s at line %d already exists", strings.Fields(data[i])[0], i+1)
			}
			//otherwise add it with empty value
			MyMap[&Vertice{Name: strings.Fields(data[i])[0]}] = struct{}{}
		}
		// if line == Start or End
		if data[i] == Start || data[i] == End {
			if len(strings.Fields(data[i+1])) != 3 {
				log.Fatalf("Wrong format of inputted room: %s", (data[i+1]))
			}
			Room := Vertice{Name: strings.Fields(data[i+1])[0]}
			if data[i] == Start {
				Room.Start = true
			} else {
				Room.End = true
			}
			if _, exists := MyMap[&Room]; exists {
				log.Fatalf("Sorry, but room %s at line %d already exists", Room.Name, i+2)
			}
			MyMap[&Room] = struct{}{}
			i++
		}
		//if 1st line of links block detected
		if strings.Contains(data[i], "-") && len(strings.Fields(data[i])) == 1 {
			index = i
			break
		}
	}

	return MyMap, index
}

func ValidateLinks(data []string, MyMap *Map) Map {
	// Check if all the called rooms exist
	allLinks := make(map[string][]string)
	links := make(map[string]struct{})

	for _, item := range data {
		temp := strings.Split(item, "-")
		if temp[0] == temp[1] {
			log.Fatal("Room linking to itself")
		}
		if _, ok := allLinks[temp[0]]; ok {
			for _, v := range allLinks[temp[0]] {
				if v == temp[1] {
					log.Fatalf("Found duplicate %s linked to room %s", v, temp[0])
				}
			}
		}
		links[temp[0]] = struct{}{}
		links[temp[1]] = struct{}{}
		allLinks[temp[0]] = append(allLinks[temp[0]], temp[1])
	}

	// Check if all the rooms from the links block
	// match the rooms collected from the rooms block
	if len(*MyMap) != len(links) {
		log.Fatal("Inexistant or missing room(s) detected within the block of links")
	}

	for k := range *MyMap {
		if _, ok := links[k.Name]; !ok {
			log.Fatalf("Room %s not found", k.Name)
		}
		items := allLinks[k.Name]
		LinksBinder(k, items, MyMap)
	}
	return *MyMap
}

func LinksBinder(Key *Vertice, items []string, MyMap *Map) {

	for _, v := range items {
		for k := range *MyMap {
			if k.Name == v {
				Key.Links = append(Key.Links, k)
				k.Links = append(k.Links, Key)
				break
			}
		}
	}
}
