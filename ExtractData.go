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

func ValidateLinks(Data []string, MyMap Map) Map {
	//Check if all the called rooms exists
	var Links string
	AllLinks := make(map[string][]string)

	for _, item := range Data {
		//separate the links by "-"
		temp := strings.Split(item, "-")
		value, exists := AllLinks[temp[0]]
		Duplicate := strings.Contains(strings.Join(value, ""), temp[1])

		//the variable Links instantiates all unique rooms
		//called within its respective block
		if !strings.Contains(Links, temp[0]) {
			Links += temp[0]
		} else if Links != "" && !strings.Contains(Links, temp[1]) {
			Links += temp[1]
		}
		//check whether the line contains a duplicated link to a specific room
		//check if the room is linking to itself
		if temp[0] == temp[1] || (exists && Duplicate) {
			log.Fatal("Room linking to itself or Duplicate link to room found")
		} else if !exists {
			AllLinks[temp[0]] = []string{temp[1]}
		} else {
			value = append(value, temp[1])
		}

	}

	for k := range MyMap {
		if len(MyMap) != len(Links) {
			log.Fatal("inexistant or missing room(s) have been detected within the block of links")
		} else if strings.Contains(Links, k.Name) && len(MyMap) == len(Links) {
			continue
		} else {
			log.Fatalf("Room %s Not found", k.Name)
		}
	}
	// DEFINITELY REQUIRES A FUNCTION WHICH WILL TAKE AS ARGUMENTS:
	// K Room,
	// multidimensional For loop, ranging over MyMap and AllLinks
	// in order to bind rooms between themselves

	for k := range MyMap {
		items := AllLinks[k.Name]
		LinksBinder(k, items, MyMap)
	}
	return MyMap

}

func LinksBinder(Key Room, items []string, MyMap Map) {
	value := MyMap[Key]
	for _, v := range items {
		for k := range MyMap {
			if k.Name == v {
				value = append(value, k)
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
