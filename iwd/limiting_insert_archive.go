package iwd

// LimitingInsert is from the tutorial of moea , which uses boxes to update the best archive
func (ar *Archive) LimitingInsert(newElement *ArchiveElement) {
	newElement.calculateBox()
	isInsertRemovesBox := ar.removeDominatedBox(newElement)
	if isInsertRemovesBox {
		return
	}

	isInsertReplaceBoxElement := ar.replaceDominatedBoxElement(newElement)
	if isInsertReplaceBoxElement {
		return
	}

	ar.insertIfNewElementIsNotDominated(newElement)
}

func (ar *Archive) removeDominatedBox(newElement *ArchiveElement) bool {

	newBox := newElement.Box
	var idxToRemoveList []int
	for arIdx := range ar.ElementList {
		element := ar.ElementList[arIdx]
		box := element.Box
		if newBox.isDominate(box) {
			idxToRemoveList = append(idxToRemoveList, arIdx)
		}
	}

	if len(idxToRemoveList) == 0 {
		return false
	}

	for idx := len(idxToRemoveList) - 1; idx >= 0; idx-- {
		idxToRemove := idxToRemoveList[idx]
		ar.ElementList = append(ar.ElementList[:idxToRemove], ar.ElementList[idxToRemove+1:]...)
	}
	ar.ElementList = append(ar.ElementList, newElement)
	return true
}

func (ar *Archive) replaceDominatedBoxElement(newElement *ArchiveElement) bool {

	tolerance := newElement.Wd.Config.Tolerance
	newBox := newElement.Box
	newScore := newElement.Wd.Score
	for arIdx := range ar.ElementList {
		element := ar.ElementList[arIdx]
		box := element.Box
		score := element.Wd.Score
		if newBox.isEqual(box) && newScore.IsDominate(score, tolerance) {
			ar.ElementList[arIdx] = newElement
			return true
		}
	}
	return false
}

func (ar *Archive) insertIfNewElementIsNotDominated(newElement *ArchiveElement) {

	tolerance := newElement.Wd.Config.Tolerance
	newBox := newElement.Box
	newScore := newElement.Wd.Score
	for arIdx := range ar.ElementList {
		element := ar.ElementList[arIdx]
		box := element.Box
		score := element.Wd.Score
		if newBox.isEqual(box) || score.IsDominate(newScore, tolerance) {
			return
		}
	}

	ar.ElementList = append(ar.ElementList, newElement)
}
