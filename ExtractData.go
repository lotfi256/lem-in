package lemin

import (
	"log"
	"strconv"
	"strings"
)

var Start, End string = "##start", "##end"

type Room struct {
	Name  string
	Start bool
	End   bool
}

type Map map[Room][]Room

func ValidateAnts(data []string) int {
	NumLine1, _ := strconv.Atoi(data[0])
	NumLine2, _ := strconv.Atoi(data[1])
	if NumLine2 != 0 {
		log.Fatalf("Invalid format of inputted ants at line 2")
	} else if NumLine1 == 0 {
		log.Fatalf("Not enough ants input to validate")
	}
	return NumLine1
}

// Haven't handle '#' or 'L' Yet.
// Did handle duplicates and Room which starts and ends.
// Initiate a counter for start and end through a different function.
func ValidateRooms(data []string) (Map, int) {
	MyMap := make(Map)
	var index int

	for i := 1; i < len(data); i++ {

		room := strings.Fields(data[i])
		if len(room) == 3 {
			Rooms := Room{Name: room[0]}
			if _, exists := MyMap[Rooms]; exists {
				log.Fatalf("Sorry, but room %s at line %d already exists", Rooms.Name, i+1)
			}
			MyMap[Rooms] = nil
		} else if data[i] == Start || data[i] == End {
			if len(strings.Fields(data[i+1])) != 3 {
				log.Fatalf("Wrong format of inputted room: %s", (data[i+1]))
			}
			Rooms := Room{Name: strings.Fields(data[i+1])[0]}
			if data[i] == Start {
				Rooms.Start = true
			} else {
				Rooms.End = true
			}
			if _, exists := MyMap[Rooms]; exists {
				log.Fatalf("Sorry, but room %s at line %d already exists", Rooms.Name, i+2)
			}
			MyMap[Rooms] = nil
			i++
		}
		if strings.Contains(data[i], "-") && len(strings.Fields(data[i])) == 1 {
			index = i
			break
		}
	}

	return MyMap, index
}

func ValidateLinks(data []string, myMap Map) Map {
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
					log.Fatal("Duplicate link to room found")
				}
			}
		}
		links[temp[0]] = struct{}{}
		links[temp[1]] = struct{}{}
		allLinks[temp[0]] = append(allLinks[temp[0]], temp[1])
	}

	// Check if all the rooms from the links block
	// match the rooms collected from the rooms block
	if len(myMap) != len(links) {
		log.Fatal("Inexistant or missing room(s) detected within the block of links")
	}

	for k := range myMap {
		if _, ok := links[k.Name]; !ok {
			log.Fatalf("Room %s not found", k.Name)
		}
		items := allLinks[k.Name]
		LinksBinder(k, items, myMap)
	}
	return myMap
}

func LinksBinder(Key Room, items []string, MyMap Map) {
	for _, v := range items {
		for k := range MyMap {
			if k.Name == v {
				MyMap[Key] = append(MyMap[Key], k)
				break
			}
		}
	}
}

// // STEP 2:
// func Validaterooms(Data []string, Starter int) (map[string][]string, int, int) {
// 	var result map[string][]string
// 	// Initialize a counter for ##start and ##end
// 	Scounter, Ecounter := 0, 0
// 	Sindex, Eindex := 0, 0
// 	for i, v := range Data[Starter:] {
// 		switch v {
// 		case Start:
// 			Scounter++
// 			Sindex = i
// 		case End:
// 			Ecounter++
// 			Eindex = i
// 		}
// 	}

// else {
// 		for _, v := range Data[Sindex : Eindex+2] {
// 			result[v] = nil
// 		}
// 		//What if out of range?
// 	}
// 	return result, Sindex, Eindex
// }
