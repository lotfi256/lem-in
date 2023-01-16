package lemin

import (
	"sync"
)

var wg sync.WaitGroup

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
			Result.Tail, Result.Head = key, key
		}
		if key.End {
			key.Parents = key.Children
			key.Children = nil
		}
	}
	LinkedList(Result.Head)
}

// Should implement the level system (increment it by 1 for each edge between it and start)
func LinkedList(Node *Vertice) {
	for _, v := range Node.Children {

		if v.End {
			continue
		}
		//add the parent node to the child
		//only if it doesn't already exist
		// for I, P := range v.Parents {
		// 	if P == Node {
		// 		break
		// 	}
		// 	if I == (len(v.Parents)-1) && P != Node {
		v.Parents = append(v.Parents, Node)
		// 	}
		// }

		//remove the Parent node from the LinkedList
		for index, item := range v.Children {
			if item.Name == Node.Name && index != len(v.Children)-1 {
				v.Children = append(v.Children[:index], v.Children[index+1:]...)
			} else if item.Name == Node.Name && index == len(v.Children)-1 {
				v.Children = v.Children[:index]
			}
		}
		Result.Head = v
		wg.Add(1)
		go LinkedList(Result.Head)
		//repeat for each child node at Result.Head
		defer wg.Done()

		// if !v.End && v.Children == nil {
		// 	delete(MyMap[v])
		// }
	}
}

//STEP 2: initialize a BFS or DFS (pick the quickest method)
//STEP 3: queue the ants using edmonds-karp method

//Variables to use: Rooms []string, Links []string, NumberOfAnts
//The above information is meant to define the map and the movements

//A function to create to find all unique paths.
//BFS or DFS?

//A function to sort out all *VALID* paths from shortest to longest.
