package iwd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/order"
)

var configTest = &config.Config{
	DeliveryDate:      "2018-08-16",
	MaxDriverCapacity: 5,
	MaxDriverDistance: 30,
	MaxTreeEntry:      4,
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

var orderTest = []*order.Order{
	&order.Order{ID: "1", UserID: "1", Quantity: 3, Coordinate: &coordinate.Coordinate{Latitude: 2, Longitude: 2}},
	&order.Order{ID: "2", UserID: "2", Quantity: 2, Coordinate: &coordinate.Coordinate{Latitude: 6, Longitude: 5}},
	&order.Order{ID: "3", UserID: "3", Quantity: 3, Coordinate: &coordinate.Coordinate{Latitude: 3, Longitude: 7}},
	&order.Order{ID: "4", UserID: "4", Quantity: 2, Coordinate: &coordinate.Coordinate{Latitude: 4, Longitude: 10}},
	&order.Order{ID: "5", UserID: "5", Quantity: 3, Coordinate: &coordinate.Coordinate{Latitude: 9, Longitude: 2}},
	&order.Order{ID: "6", UserID: "6", Quantity: 2, Coordinate: &coordinate.Coordinate{Latitude: 9, Longitude: 8}},
	&order.Order{ID: "7", UserID: "7", Quantity: 3, Coordinate: &coordinate.Coordinate{Latitude: 6, Longitude: 11}},
	&order.Order{ID: "8", UserID: "8", Quantity: 2, Coordinate: &coordinate.Coordinate{Latitude: 7, Longitude: 7}},
	&order.Order{ID: "9", UserID: "9", Quantity: 3, Coordinate: &coordinate.Coordinate{Latitude: 6, Longitude: 2}},
	&order.Order{ID: "10", UserID: "10", Quantity: 2, Coordinate: &coordinate.Coordinate{Latitude: 3, Longitude: 4}},
}

var kitchenTest = []*kitchen.Kitchen{
	&kitchen.Kitchen{
		ID: "1",
		Capacity: &kitchen.Capacity{
			Minimum: 4,
			Optimum: 8,
			Maximum: 12,
		},
		Preference: 0.,
		Coordinate: &coordinate.Coordinate{
			Latitude:  8,
			Longitude: 4,
		},
	},
	&kitchen.Kitchen{
		ID: "2",
		Capacity: &kitchen.Capacity{
			Minimum: 4,
			Optimum: 8,
			Maximum: 12,
		},
		Preference: 0.,
		Coordinate: &coordinate.Coordinate{
			Latitude:  5,
			Longitude: 8,
		},
	},
	&kitchen.Kitchen{
		ID: "3",
		Capacity: &kitchen.Capacity{
			Minimum: 4,
			Optimum: 8,
			Maximum: 12,
		},
		Preference: 0.,
		Coordinate: &coordinate.Coordinate{
			Latitude:  2,
			Longitude: 3,
		},
	},
}

var soilMapTC = struct {
	kitchen      *kitchen.Kitchen
	order        *order.Order
	expectedSoil float64
	newSoil      float64
}{
	kitchen:      kitchenTest[1],
	order:        orderTest[2],
	expectedSoil: 1000,
	newSoil:      200,
}

func TestSoilMap(t *testing.T) {

	// Arrange
	soilMap := NewSoilMap(kitchenTest, orderTest, configTest)
	kitchen := soilMapTC.kitchen
	order := soilMapTC.order

	// Act
	soil := soilMap.GetSoil(kitchen, order)
	expectedSoil := soilMapTC.expectedSoil
	soilMap.UpdateSoil(kitchen, order, soilMapTC.newSoil)
	newSoil := soilMap.GetSoil(kitchen, order)
	expectedNewSoil := soilMapTC.newSoil

	// Assert
	t.Run("Test Get Soil", func(t *testing.T) {
		assert.Equalf(t, expectedSoil, soil, "Error!expected soil %f, got %f", expectedSoil, soil)
	})
	t.Run("Test Update Soil", func(t *testing.T) {
		assert.Equalf(t, expectedNewSoil, newSoil, "Error! expected soil after update %f, got %f", expectedNewSoil, newSoil)
	})

}

var ksqTC = struct {
	kitchen *kitchen.Kitchen
	addQty  int
}{
	kitchen: kitchenTest[0],
	addQty:  10,
}

func TestKitchenServedQtyMap(t *testing.T) {
	// Arrange
	ksqMap := NewKitchenServedQtyMap(kitchenTest)
	kitchen := ksqTC.kitchen

	// Act
	expectedQty := 0
	qty := ksqMap.GetServedQty(kitchen)
	ksqMap.AddQty(kitchen, ksqTC.addQty)
	newQty := ksqMap.GetServedQty(kitchen)
	expectedNewQty := ksqTC.addQty

	// Assert
	t.Run("Test Get Served Qty", func(t *testing.T) {
		assert.Equalf(t, expectedQty, qty, "Error! expected served qty %d, got %d", expectedQty, qty)
	})
	t.Run("Test Add Qty", func(t *testing.T) {
		assert.Equalf(t, expectedNewQty, newQty, "Error! expected qty after addition %d, got %d", expectedNewQty, newQty)
	})

}
