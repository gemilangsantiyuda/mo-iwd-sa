package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vroup/mo-iwd-sa/config"
)

var testCase = config.Config{
	DataSize:          "2018-08-16",
	MaxDriverCapacity: 40,
	MaxDriverDistance: 90000,
	MaxTreeEntry:      60,
	IwdParameter: config.IwdParameter{
		MaximumIteration: 100,
		PopulationSize:   10,
		As:               1000,
		Bs:               0.01,
		Cs:               1,
		Av:               1000,
		Bv:               0.01,
		Cv:               1,
		InitSoil:         1000,
		InitIWDVel:       100,
	},
	Weight: config.Weight{
		RiderCost:         0.7,
		KitchenOptimality: 0.2,
		UserSatisfaction:  0.1,
	},
}

func TestConfig(t *testing.T) {
	// Arrange
	var conf *config.Config

	// Act
	conf = config.ReadConfig()

	// assert
	assert.Equalf(t, testCase, conf, "Error! expected %+v, got %+v", testCase, conf)
}
