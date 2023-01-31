package lemin

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
)

var AllPaths [][]Vertice

type Vertice struct {
	Name    string
	Start   bool
	End     bool
	Links   []*Vertice
	Vacant  bool
	Visited bool
}

// STEP 1: Get all possible combinations of paths
func RecursivePathFinder(Node *Vertice, route []Vertice) {
	if Node.End {
		route = append(route, *Node)
		sepRoute := make([]Vertice, len(route))
		copy(sepRoute, route)
		AllPaths = append(AllPaths, sepRoute)
		return
	} else {
		if inArray(route, *Node) {
			return
		}
		route = append(route, *Node)
	}

	for _, v := range Node.Links {
		RecursivePathFinder(v, route)
	}
}

// Hello world
// find all combinations of unique paths
func CombinePaths(AllPaths [][]Vertice) [][][]Vertice {
	//THE UGLIEST FUNCTION I HAVE EVER WRITTEN
	if len(AllPaths) == 0 {
		log.Fatal("no path found from start to end")
	}

	Result := make([][][]Vertice, 0)
	CombPaths := make([][]Vertice, 0)
	var counter int
	var Breaker bool

	for _, P1 := range AllPaths {
		//compare 1 Path with all other Paths
		CombPaths = append(CombPaths, P1)
		for _, P2 := range AllPaths {
			//compare the Paths node by node.
			for i, P := range CombPaths {
				if !Breaker {
					for _, v := range P[1 : len(P)-1] {
						if inArray(P2[1:len(P2)-1], v) {
							Breaker = true
							break
						}
					}
				}
				if i == len(CombPaths)-1 && !Breaker {
					CombPaths = append(CombPaths, P2)
				}
			}
			Breaker = false
		}
		if counter <= len(CombPaths) {
			Result = append(Result, CombPaths)
			counter = len(CombPaths)
		}

		CombPaths = nil
	}
	return Result
}

// STEP 2: Find the Maximum Flow
func ChoosePath(CombPaths [][][]Vertice) [][]Vertice {
	var Max int
	var Sum int
	var Index int

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
	temp := 0
	for I, P := range CombPaths {

		if len(P) == Max {
			for i, path := range P {

				Sum += len(path)
				if i == len(P)-1 {
					if temp == 0 {
						Index = I
						temp = Sum
					} else if Sum <= temp {
						Index = I
						temp = Sum
					}
					Sum = 0

				}
			}
		}

	}
	return CombPaths[Index]
}

// STEP 3: queue the ants using edmonds-karp method
func QueueThem(NumAnts int, MaxFlow [][]Vertice) {
	//Sort them from shortest to longest
	sort.Slice(MaxFlow, func(i, j int) bool { return len(MaxFlow[j]) > len(MaxFlow[i]) })

	//start queuing them using edmonds-karp
	QueuedAnts := make([][]string, len(MaxFlow))

	//here, we are adding all ants to the only path we have
	//hence why len(MaxFlow) would be 1
	if len(MaxFlow) == 1 {
		for i := 1; i <= NumAnts; i++ {
			AntName := "L" + strconv.Itoa(i)
			QueuedAnts[0] = append(QueuedAnts[0], AntName)
		}
	} else {
		for i := 1; i <= NumAnts; i++ {
			AntName := "L" + strconv.Itoa(i)
			//after adding an ant to the queue
			//we need to decide which path does it
			//correspond to
			for j := 0; j < len(MaxFlow); j++ {
				if j < len(MaxFlow)-1 {
					PathSize1 := len(MaxFlow[j]) + len(QueuedAnts[j])
					PathSize2 := len(MaxFlow[j+1]) + len(QueuedAnts[j+1])
					if PathSize1 <= PathSize2 {
						QueuedAnts[j] = append(QueuedAnts[j], AntName)
						break
					}
				} else if j == len(MaxFlow)-1 {
					QueuedAnts[j] = append(QueuedAnts[j], AntName)
				}

			}
		}

	}
	for i, v := range QueuedAnts {
		fmt.Println("Path ", i+1, ": ", v)
	}

}

func PrintResult(QueuedAnts [][]string, MaxFlow [][]Vertice) {
	// func printResult(ants []Ant) {
	// 	for i := 0; i < len(ants[0].positions); i++ {
	// 		line := ""
	// 		for j := 0; j < len(ants); j++ {
	// 			line += "L" + strconv.Itoa(j+1) + "-" + strconv.Itoa(ants[j].positions[i]) + " "
	// 		}
	// 		fmt.Println(strings.TrimRight(line, " "))
	// 	}
	}

	//each ant must travel from
	//start to end simultaneously

	//Print Turn 1:
	//Print Turn 2:
	//...
	//...
	//Print Turn N:

	// Should maybe create a map[string]int
	// string represent the ants as keys
	// int to keep track of their pathing until end is reached

	// initialize the ants from the queue one by one,
	//([]string and delete elements that reached the end)
	// by incrementing the index from QueuedAnts

	// Maybe variadic function?
	// at each For-loop, add an ant to the starting point

	// for _, Path := range MaxFlow {
	// 	//Recursive for each ant L inside QueuedAnts until all reach Node.End?
	// }

}

func inArray(s []Vertice, vp Vertice) (result bool) {
	for _, v := range s {
		if reflect.DeepEqual(v, vp) {
			result = true
		}
	}
	return
}
