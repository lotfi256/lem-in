package lemin

//Maybe will use an array of []struct
//to store collected rooms and link them together
//with a []*Room field

type Room struct {
	Name   string
	Next   []*Room
	Used   bool
	Vacant bool
	Start  bool
	End    bool
	level  int
}

// STEP 1: draw the colony from MyMap
// I will extract and organize the data from the map:
func DrawMap(MyMap *Map) []*Room {
	Colony := make([]*Room, 0, len(*MyMap))

	return Colony
}

//STEP 2: initialize a BFS or DFS (pick the quickest method)
//STEP 3: queue the ants using edmonds-karp method

//Variables to use: Rooms []string, Links []string, NumberOfAnts
//The above information is meant to define the map and the movements

//A function to create to find all unique paths.
//BFS or DFS?

//A function to sort out all *VALID* paths from shortest to longest.
