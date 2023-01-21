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

// STEP 1: Get all possible paths
func RecursivePathFinder(Node *Vertice, route []Vertice) {
	//to be reinvented
	//TRY TO DEBUG
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

// find all combinations of unique paths
func CombinePaths(AllPaths [][]Vertice) [][][]Vertice {
	//THE UGLIEST FUNCTION I HAVE EVER CREATED

	Result := make([][][]Vertice, 0)
	MaxFlow := make([][]Vertice, 0)
	var counter int
	var Breaker bool = false

	for _, P1 := range AllPaths {
		//compare 1 Path with all other Paths
		MaxFlow = append(MaxFlow, P1)
		for _, P2 := range AllPaths {
			//compare the Paths node by node.
			for i, P := range MaxFlow {
				if !Breaker {
					for _, v := range P[1 : len(P)-1] {
						if inArray(P2[1:len(P2)-1], v) {
							Breaker = true
							break
						}
					}
				}
				if i == len(MaxFlow)-1 && !Breaker {
					MaxFlow = append(MaxFlow, P2)
				}
			}
			Breaker = false
		}
		if counter <= len(MaxFlow) {
			Result = append(Result, MaxFlow)
			counter = len(MaxFlow)
		}

		MaxFlow = nil
	}
	return Result
}

func ChoosePath(CombPaths [][][]Vertice) [][]Vertice {
	var Max int
	var Sum int
	var Index int
	//Pick the highest max flow

	//Start by finding the highest amount of flow
	for i, j := 0, len(CombPaths)-1; i < j; i, j = i+1, j-1 {
		if len(CombPaths[i]) >= len(CombPaths[j]) && Max < len(CombPaths[i]) {
			Max = len(CombPaths[i])
		} else if Max < len(CombPaths[j]) {
			Max = len(CombPaths[j])
		}
	}
	//If several paths share the same amount of flow
	//then choose the shortest one
	for I, P := range CombPaths {
		temp := 0
		if len(P) == Max {
			for i, path := range P {
				temp += len(path)
				if i == len(P)-1 && temp > Sum {
					Index = I
				}
			}
		}
	}

	return CombPaths[Index]
}

//STEP 3: queue the ants using edmonds-karp method

func inArray(s []Vertice, vp Vertice) (result bool) {
	for _, v := range s {
		if reflect.DeepEqual(v, vp) {
			result = true
		}
	}
	return
}
