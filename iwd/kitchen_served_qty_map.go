package iwd

// import "github.com/vroup/mo-iwd-sa/kitchen"

// // KitchenServedQtyMap map the kitchen and the served qunatity so far
// type KitchenServedQtyMap map[string]int

// // NewKitchenServedQtyMap make a new map from kitchenlist
// func NewKitchenServedQtyMap(kitchenList []*kitchen.Kitchen) KitchenServedQtyMap {
// 	ksqMap := make(KitchenServedQtyMap)
// 	for idx := range kitchenList {
// 		key := kitchenList[idx].GetID()
// 		ksqMap[key] = 0
// 	}
// 	return ksqMap
// }

// // GetServedQty return the kitchen's served qty
// func (ksqMap KitchenServedQtyMap) GetServedQty(kitchen *kitchen.Kitchen) int {
// 	key := kitchen.GetID()
// 	return ksqMap[key]
// }

// // AddQty add some quantit to kitchen served Qty
// func (ksqMap KitchenServedQtyMap) AddQty(kitchen *kitchen.Kitchen, addQty int) {
// 	key := kitchen.GetID()
// 	ksqMap[key] += addQty
// }
