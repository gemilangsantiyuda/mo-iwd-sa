package iwd

import (
	"math"

	"github.com/vroup/mo-iwd-sa/config"
)

// func getSAProb(score1, score2 *Score, config *config.Config) float64 {
// 	temp := config.SaParam.Temp
// 	// fmt.Println(temp)
// 	rdif := score1.RiderCost - score2.RiderCost
// 	rdProb := math.Exp(-float64(rdif) / temp.RiderCost)
// 	// fmt.Println("RiderCost Difference ", rdif, "prob ", rdProb)

// 	kdif := score1.KitchenOptimality - score2.KitchenOptimality
// 	kdProb := math.Min(1, math.Exp(-float64(kdif)/temp.KitchenOptimality))
// 	// fmt.Println("Kitchen Optimality Difference", kdif, " prob ", kdProb)

// 	udif := score1.UserSatisfaction - score2.UserSatisfaction
// 	udProb := math.Min(1, math.Exp(-(udif)/temp.UserSatisfaction))
// 	// fmt.Println("User Rating Difference ", udif, " prob ", udProb)

// 	w := config.Weight
// 	totalProb := rdProb*w.RiderCost + kdProb*w.KitchenOptimality + udProb*w.UserSatisfaction
// 	// fmt.Println("Total Prob:", totalProb)
// 	return totalProb
// }

func getSAProb(config *config.Config, fitness float64) float64 {
	temp := config.SaParam.Temp
	prob := math.Exp(-fitness / temp)
	
	return prob
}
