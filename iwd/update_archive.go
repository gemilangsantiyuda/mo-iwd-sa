package iwd

import (
	"sort"
	"sync"
)

// Update according to SPEAII
func (ar *Archive) Update(maxSize int) {

	if len(ar.ElementList) < 2 {
		return
	}
	if maxSize > len(ar.ElementList) {
		maxSize = len(ar.ElementList)
	}

	SFit := make([]float64, len(ar.ElementList))
	for idx1 := range ar.ElementList {
		wd1 := ar.ElementList[idx1].Wd
		for idx2 := range ar.ElementList {
			if idx1 != idx2 {
				wd2 := ar.ElementList[idx2].Wd
				if wd1.Score.IsDominate(wd2.Score, wd1.Config.Tolerance) {
					SFit[idx1] += 1.
				}
			}
		}
	}

	RFit := make([]float64, len(ar.ElementList))
	for idx1 := range ar.ElementList {
		wd1 := ar.ElementList[idx1].Wd
		for idx2 := range ar.ElementList {
			if idx1 != idx2 {
				wd2 := ar.ElementList[idx2].Wd
				if wd2.Score.IsDominate(wd1.Score, wd1.Config.Tolerance) {
					RFit[idx1] += SFit[idx2]
				}
			}
		}
	}

	diffMatrix := calcDiffMatrix(ar)
	for idx := range ar.ElementList {
		ar.ElementList[idx].Fitness = RFit[idx] + 1/(diffMatrix[idx][1]+2)
	}

	sort.Slice(ar.ElementList, func(i, j int) bool {
		return ar.ElementList[i].Fitness < ar.ElementList[j].Fitness
	})

	// remove the wd with fitness > 1 , if less then maxsize got then add the remaining until = maxsize
	badIdx := -1
	for idx := range ar.ElementList {
		if ar.ElementList[idx].Fitness >= 1 {
			badIdx = idx
			break
		}
	}
	if badIdx != -1 {
		if badIdx < maxSize {
			badIdx = maxSize
		}
		ar.ElementList = ar.ElementList[:badIdx]
	}

	for len(ar.ElementList) > maxSize {
		diffMatrix = calcDiffMatrix(ar)
		removedIdx := getElementIdxToRemove(ar, diffMatrix)
		ar.ElementList = append(ar.ElementList[:removedIdx], ar.ElementList[removedIdx+1:]...)
	}

}

func getElementIdxToRemove(archive *Archive, diffMatrix [][]float64) int {
	chosenIdx := -1
	for idx1 := range archive.ElementList {
		for idx2 := range archive.ElementList {
			hasSmallerDiff := true
			for idx := range archive.ElementList {
				if diffMatrix[idx1][idx] < diffMatrix[idx2][idx] {
					break
				}
				if diffMatrix[idx1][idx] > diffMatrix[idx2][idx] {
					hasSmallerDiff = false
					break
				}
			}
			if !hasSmallerDiff {
				break
			}
			if idx2 == len(archive.ElementList)-1 {
				chosenIdx = idx1
			}
		}
		if chosenIdx > -1 {
			break
		}
	}

	return chosenIdx
}

func calcDiffMatrix(archive *Archive) [][]float64 {
	wg := sync.WaitGroup{}
	diffMatrix := make([][]float64, 0)
	for idx := range archive.ElementList {
		diffList := make([]float64, len(archive.ElementList))
		diffMatrix = append(diffMatrix, diffList)
		wg.Add(1)
		go calcDiffList(&wg, archive, idx, diffMatrix[idx])
	}
	wg.Wait()
	return diffMatrix
}

func calcDiffList(wg *sync.WaitGroup, archive *Archive, currentIdx int, diffList []float64) {
	defer wg.Done()

	currentScore := archive.ElementList[currentIdx].Wd.Score
	for idx := range archive.ElementList {
		score := archive.ElementList[idx].Wd.Score
		diff := currentScore.GetDifference(score)
		diffList[idx] = diff
	}

	sort.Slice(diffList, func(i, j int) bool {
		return diffList[i] < diffList[j]
	})
}
