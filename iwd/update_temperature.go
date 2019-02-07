package iwd

import "github.com/vroup/mo-iwd-sa/config"

// func updateTemperature(config *config.Config) {
// 	config.SaParam.Temp.RiderCost *= config.SaParam.CoolingRate
// 	config.SaParam.Temp.KitchenOptimality *= config.SaParam.CoolingRate
// 	config.SaParam.Temp.UserSatisfaction *= config.SaParam.CoolingRate
// }

func updateTemperature(config *config.Config) {
	config.SaParam.Temp *= config.SaParam.CoolingRate
}
