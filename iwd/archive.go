package iwd

import (
	"math"

	"github.com/vroup/mo-iwd-sa/config"
)

// ArchiveElement for the archive
type ArchiveElement struct {
	Wd      *WaterDrop
	Fitness float64
	Box     *Box
}

// func (ae *ArchiveElement) calculateBox() {
// 	score := ae.Wd.Score
// 	tolerance := ae.Wd.Config.Tolerance
// 	rcCeil := math.Ceil(float64(score.RiderCost) / float64(tolerance.RiderCost))
// 	koCeil := math.Ceil(float64(score.KitchenOptimality) / float64(tolerance.KitchenOptimality))
// 	usCeil := math.Ceil(score.UserSatisfaction / tolerance.UserSatisfaction)
// 	ae.Box = &Box{
// 		RiderCost:         rcCeil,
// 		KitchenOptimality: koCeil,
// 		UserSatisfaction:  usCeil,
// 	}
// }

//Box is the ceil boundary of the score of a solution if divided by each of their tolerance,
type Box struct {
	RiderCost         float64
	KitchenOptimality float64
	UserSatisfaction  float64
}

func (box *Box) isDominate(box2 *Box) bool {
	if (box.RiderCost <= box2.RiderCost) && (box.KitchenOptimality <= box2.KitchenOptimality) && (box.UserSatisfaction <= box2.UserSatisfaction) {
		return (box.RiderCost < box2.RiderCost) || (box.KitchenOptimality < box2.KitchenOptimality) || (box.UserSatisfaction < box2.UserSatisfaction)
	}
	return false
}

func (box *Box) isEqual(box2 *Box) bool {
	return (box.RiderCost == box2.RiderCost) && (box.KitchenOptimality == box2.KitchenOptimality) && (box.UserSatisfaction == box2.UserSatisfaction)
}

// Archive for solutions,, adapting SPEAII
type Archive struct {
	ElementList []*ArchiveElement
}

func (ar *Archive) getIGD(referencePoint []*Score, conf *config.Config) float64 {
	igd := 0.
	for idx := range referencePoint {
		minDif := math.Inf(1)
		score1 := referencePoint[idx]
		for idx2 := range ar.ElementList {
			score2 := ar.ElementList[idx2].Wd.Score
			diff := score1.GetDifference(score2, conf)
			minDif = math.Min(minDif, diff)
		}
		igd += minDif
	}
	igd /= float64(len(referencePoint))
	return igd
}
