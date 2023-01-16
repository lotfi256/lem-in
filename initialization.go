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

// STEP 1: draw the colony from MyMap
// I will extract and organize the data from the map:
func AdjustMap(MyMap *Map) {

	//Starting room is the first entry in the LinkedList
	for key := range *MyMap {
		if key.Start {
			//the starting point will always be Tail
			//Head will move across the map
			Result.Tail = key
		}
		if key.End {
			Result.Head = key
			key.Parents = key.Children
			key.Children = nil
		}
	}
	for Result.Tail.Children != nil {
		LinkedList(Result.Tail)
	}
}

// maybe use a go routine
// To wait for each function to finish

// or maybe loop over the unvisited nodes
func LinkedList(Node *Vertice) {
	Node.Visited = true
	for _, v := range Node.Children {
		if !v.End && !v.Visited {
			v.Parents = append(v.Parents, Node)
			for index, item := range v.Children {
				if item.Name == Node.Name {
					v.Children = append(v.Children[:index], v.Children[index+1:]...)
				}
			}
		} else {
			break
		}
		LinkedList(v)
	}
}

//STEP 2: initialize a BFS or DFS (pick the quickest method)
//STEP 3: queue the ants using edmonds-karp method

//Variables to use: Rooms []string, Links []string, NumberOfAnts
//The above information is meant to define the map and the movements

//A function to create to find all unique paths.
//BFS or DFS?

//A function to sort out all *VALID* paths from shortest to longest.
