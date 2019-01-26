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
	for _ = range gsTestCase.valueList {
		idx := chooseIdxByExpRank(gsTestCase.valueList, 0.6)
		fmt.Println(idx)
	}

	assert.Equal(t, 0, 1, "")
}
