package iwd

import (
	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
)

// Solve the mdovrp returning the best waterdrop
func Solve(orderList []*order.Order, kitchenList *[]kitchen.Kitchen, tree *mtree.Tree, config *config.Config) *WaterDrop {

	return nil
}
