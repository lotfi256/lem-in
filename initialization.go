package lemin

// Maybe will use an array of []struct
// to store collected rooms and link them together
// with a []*Vertice field
type Vertice struct {
	Name         string
	Start        bool
	End          bool
	Children     []*Vertice
	Parents      []*Vertice
	Vacant       bool
	Visited      bool
	CurrentLevel int
}

type LL struct {
	Head *Vertice
	Tail *Vertice
}

var Result LL

func AdjustMap(MyMap *Map) {

	//Starting room is the first entry in the LinkedList
	for key := range *MyMap {
		if key.Start {
			//the starting point will always be Tail
			//Head will move across the map
			Result.Head, Result.Tail = key, key
		}
		if key.End {
			key.Parents = key.Children
			key.Children = nil
		}
	}
	for Result.Head.Children != nil {
		LinkedList(Result.Head)
	}
}

// maybe use a go routine
// To wait for each function to finish

// or maybe loop over the unvisited nodes
func LinkedList(Node *Vertice) {
	Node.Visited = true
	for _, v := range Node.Children {
		if !v.End {
			v.Parents = append(v.Parents, Node)
			for index, item := range v.Children {
				if item.Name == Node.Name {
					v.Children = append(v.Children[:index], v.Children[index+1:]...)
				}
			}
			Result.Head = v
		}
	}
}

// //THIS SHOULD BE A RECURSIVE FUNCTION
// for _, room := range Result.Head.Children {
// 	room.Parents = append(room.Parents, Result.Head)
// 	//End room has no children
// 	if room.End {
// 		room.Parents = room.Children
// 		room.Children = nil
// 	}
// 	//remove Parent nodes from children
// 	for i, element := range room.Children {
// 		if Result.Head == element && i < len(room.Children) {
// 			room.Children = append(room.Children[:i], room.Children[i+1:]...)
// 		} else if Result.Head == element {
// 			room.Children = room.Children[:i]
// 		}
// 	}

// }

// STEP 1: draw the colony from MyMap
// I will extract and organize the data from the map:
func DrawMap(linkedlist *LL) {
	// will link from
}

//STEP 2: initialize a BFS or DFS (pick the quickest method)
//STEP 3: queue the ants using edmonds-karp method

//Variables to use: Rooms []string, Links []string, NumberOfAnts
//The above information is meant to define the map and the movements

//A function to create to find all unique paths.
//BFS or DFS?

//A function to sort out all *VALID* paths from shortest to longest.
