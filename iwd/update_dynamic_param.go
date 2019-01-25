package iwd

import "math"

func (wd *WaterDrop) updateDynamicParameter(currentNode, nextNode node, moDistance float64) {
	wd.updateVelocity(currentNode, nextNode)
	dSoil := wd.getDeltaSoil(currentNode, nextNode, moDistance)
	wd.updateEdgeSoil(currentNode, nextNode, dSoil)
	wd.Soil += dSoil
}

func (wd *WaterDrop) getDeltaSoil(currentNode, nextNode node, moDistance float64) float64 {
	iwdParam := wd.Config.IwdParameter
	t := moDistance / math.Max(0.0001, wd.Velocity)
	dSoil := iwdParam.As / (iwdParam.Bs + (iwdParam.Cs * t))
	return dSoil
}

func (wd *WaterDrop) updateEdgeSoil(currentNode, nextNode node, dSoil float64) {
	soilMap := wd.SoilMap
	p := wd.Config.IwdParameter.P
	edgeSoil := soilMap.GetSoil(currentNode, nextNode)
	newEdgeSoil := (1-p)*edgeSoil - p*dSoil
	soilMap.UpdateSoil(currentNode, nextNode, newEdgeSoil)

}

func (wd *WaterDrop) updateVelocity(currentNode, nextNode node) {
	iwdParam := wd.Config.IwdParameter
	soilMap := wd.SoilMap
	edgeSoil := soilMap.GetSoil(currentNode, nextNode)
	newVelocity := wd.Velocity + (iwdParam.Av / (iwdParam.Bv + iwdParam.Cv*edgeSoil))
	wd.Velocity = newVelocity
}
