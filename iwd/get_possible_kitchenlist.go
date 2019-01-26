package iwd

import "github.com/vroup/mo-iwd-sa/kitchen"

// Possible kitchens have at least one possible order to serve whose distance < maxDriverDistance and order.Qty < (kitchenCap.Max-kitchenServedQty)
func (wd *WaterDrop) getPossibleKitchenList() []*kitchen.Kitchen {
	kitchenList := wd.KitchenList
	var possibleKitchenList []*kitchen.Kitchen
	tree := wd.Tree
	for idx := range kitchenList {
		kitchen := kitchenList[idx]
		maxDistance := wd.Config.MaxDriverDistance
		remainingKitchenCap := kitchen.Capacity.Maximum - wd.KitchenServedQtyMap.GetServedQty(kitchen)
		maxCap := wd.Config.MaxDriverCapacity
		if maxCap > remainingKitchenCap {
			maxCap = remainingKitchenCap
		}
		neighbourList := tree.KnnSearch(tree.Root, kitchen, 1, maxCap, maxDistance)
		// if there's neighbour found then this kitchen has possibility
		if len(neighbourList) > 0 {
			possibleKitchenList = append(possibleKitchenList, kitchen)
		}
	}
	return possibleKitchenList
}
