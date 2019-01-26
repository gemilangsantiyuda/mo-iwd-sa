package iwd

func nextStepWD(workingWaterDrops, finishedWaterDrops []*WaterDrop) ([]*WaterDrop, []*WaterDrop) {

	finishWDFlags := make([]bool, len(workingWaterDrops))

	for wdIdx := range workingWaterDrops {
		wd := workingWaterDrops[wdIdx]
		// fmt.Println(wd)
		// fmt.Println(wd.Soil, wd.Velocity)
		currentRoute := wd.getCurrentRoute()
		// fmt.Println(currentRoute)
		currentNode := wd.getCurrentNode(currentRoute)
		// fmt.Println(currentNode)
		nextOrder, distance, moDistance := wd.getNextOrder(currentRoute, currentNode)
		// if no order can be visited anymore then create a new route,, and if no route can be created anymore the water drop is finished traversing
		if nextOrder == nil {
			currentRoute.CalcRiderCost(wd.Config)
			newRoute := wd.createNewRoute()
			if newRoute == nil {
				finishWDFlags[wdIdx] = true
				break
			} else {
				wd.RouteList = append(wd.RouteList, newRoute)
			}
			continue
		}
		// else we visit that next order
		wd.visitOrder(currentRoute, currentNode, nextOrder, distance, moDistance)
	}

	// move the finished waterdrop from the workingwaterdrops to finishedWaterDrops
	for idx := len(workingWaterDrops) - 1; idx >= 0; idx-- {
		if finishWDFlags[idx] {
			finishedWaterDrops = append(finishedWaterDrops, workingWaterDrops[idx])
			workingWaterDrops = append(workingWaterDrops[:idx], workingWaterDrops[idx+1:]...)
		}
	}
	return workingWaterDrops, finishedWaterDrops
}
