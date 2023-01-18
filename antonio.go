package lemin

// func LinkedList() {
// 	for _, v := range Node.Children {

// 		// if v.CurrentLevel != 0 {
// 		// // 	Node.Parents = append(Node.Parents, v)
// 		// // }

// 		//add the parent node to the child
// 		//only if it doesn't already exist

// 		if !inArray(v.Parents, Node) {
// 			v.Parents = append(v.Parents, Node)
// 			for index, item := range v.Children {
// 				if item == Node && index != len(v.Children)-1 {
// 					v.Children = append(v.Children[:index], v.Children[index+1:]...)
// 				} else if item == Node && index == len(v.Children)-1 {
// 					v.Children = v.Children[:index]
// 				}
// 			}
// 		} else {
// 			continue
// 		}

// 		//remove the Parent node from the LinkedList

// 		//Result.Head = v

// 	}
// }
// func AdjustMap(MyMap *Map) {

// 	//Starting room is the first entry in the LinkedList
// 	for key := range *MyMap {
// 		if key.Start {
// 			//the starting point will always be Tail
// 			//Head will move across the map
// 			Result.Tail, Result.Head = key, key
// 		}
// 		if key.End {
// 			key.Parents = key.Children
// 			key.Children = nil
// 		}
// 	}

// }

// // Should implement the level system (increment it by 1 for each edge between it and start)
// func LinkedList(Node *Vertice) {
// 	Node.Visited = true
// 	for i, v := range Node.Children {
// 		if v.Visited {
// 			Node.Parents = append(Node.Parents, v)
// 			Node.Children[i] = nil
// 		}
// 	}
// 	LinkedList(N)
// }
