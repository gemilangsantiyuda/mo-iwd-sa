package iwd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var gsTestCase = struct {
	valueList []float64
	m         int
}{
	valueList: []float64{1, 2, 3, 4, 5},
	m:         3,
}

func TestGS(t *testing.T) {
	// Arrange
<<<<<<< HEAD
	for _ = range gsTestCase.valueList {
=======
	for range gsTestCase.valueList {
>>>>>>> iwd-sa with SPEA 2
		idx := chooseIdxByExpRank(gsTestCase.valueList, 0.6)
		fmt.Println(idx)
	}

	assert.Equal(t, 0, 1, "")
}
