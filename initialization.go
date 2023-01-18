package lemin

import "reflect"

//var wg sync.WaitGroup

type Vertice struct {
	Name  string
	Start bool
	End   bool
	Links []*Vertice
	// Parents      []*Vertice
	Vacant       bool
	Visited      bool
	CurrentLevel int
}

var AllPaths [][]Vertice

// STEP 1: draw the colony from MyMap
// I will extract and organize the data from the map:
func RecursivePathFinder(Node *Vertice, route []Vertice) {
	if Node.End {
		route = append(route, *Node)
		AllPaths = append(AllPaths, route)
		return
	}

	if inArray(route, *Node) {
		return
	}

	route = append(route, *Node)
	for _, v := range Node.Links {
		RecursivePathFinder(v, route)
	}
}

// func PathFinder(MyMap *Map) [][]*Vertice {
// 	var result [][]*Vertice
// 	var Path []*Vertice

// }

//STEP 2: Find Maximum Flow

//find all combinations of unique paths:

func CombinePaths(AllPaths [][]Vertice) [][][]Vertice {
	Result := make([][][]Vertice, 0)
	// MaxFlow := make([][]Vertice, 0)
	// var counter int
	// // for _, Paths1 := range AllPaths {
	// // 	//compare 1 Path with all other Paths
	// // 	for _, Path2 := range AllPaths {
	// // 		//compare the Paths node by node.

	// // 	}
	// // 	if counter <= len(MaxFlow) {
	// // 		Result = append(Result, MaxFlow)
	// // 		counter = len(MaxFlow)
	// // 	}
	// // }
	return Result
}

//STEP 3: queue the ants using edmonds-karp method

// func CompareNodes(Path1 []Vertice, Path2 []Vertice) bool {
// 	k := 0
// 	for i := 1; i < len(Path1) || i < len(Path1); i++ {

// 	}
// 	return
// }

func inArray(s []Vertice, vp Vertice) (result bool) {
	for _, v := range s {
		if reflect.DeepEqual(v, vp) {
			result = true
		}
	}
	return
}
