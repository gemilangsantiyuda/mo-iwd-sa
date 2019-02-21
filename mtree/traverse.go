package mtree

import (
	"fmt"
	"log"
	"math"
)

func traverse(currentNode node) {
	// fmt.Print(currentNode.isLeaf())
	// fmt.Printf(" %v\n", currentNode)

	// entries
	// fmt.Println("ENTRIES : ")
	// entryList := currentNode.getEntryList()
	// for idx := range entryList {
	// 	entry := entryList[idx]
	// 	fmt.Printf("-------> parent:%v, radius: %f, dist : %f, centroidObject : %s \n", entry.getParent(), entry.getRadius(), entry.getDistanceFromParent(), entry.getCentroidObject().GetID())
	// }
	// fmt.Print("ENTRIES END--------------\n\n")

	if currentNode.isLeaf() {
		currentLeaf := currentNode.(*leaf)
		// recalculate radius and check the radius consistency
		recRadius := math.Inf(-1)
		for idx := range currentLeaf.entryList {
			entry := currentLeaf.entryList[idx]
			radius := entry.getRadius() + entry.getDistanceFromParent()
			if recRadius < radius {
				recRadius = radius
			}
		}

		if recRadius > currentLeaf.radius {
			log.Fatal(fmt.Sprintf("Error Recalculated Radius (%f) is bigger then the current radius (%f), node = %+v\n", recRadius, currentLeaf.getRadius(), currentLeaf))
		}

	} else {
		currentBranch := currentNode.(*branch)
		// recalculate radius and check the radius consistency
		recRadius := math.Inf(-1)
		for idx := range currentBranch.entryList {
			entry := currentBranch.entryList[idx]
			radius := entry.getRadius() + entry.getDistanceFromParent()
			if recRadius < radius {
				recRadius = radius
			}
		}

		if recRadius > currentBranch.radius {
			log.Fatal(fmt.Sprintf("Error Recalculated Radius (%f) is bigger then the current radius (%f), node = %+v\n", recRadius, currentBranch.getRadius(), currentBranch))
		}
		for idx := range currentBranch.entryList {
			nextNode := currentBranch.entryList[idx].(node)
			traverse(nextNode)
		}
	}
}
