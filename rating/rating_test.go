package rating_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

var testCase = struct {
	conf           config.Config
	userID         string
	kitchenID      string
	expectedRating float64
}{
	conf: config.Config{
		DeliveryDate: "2017-08-16",
	},
	userID:         "23589",
	kitchenID:      "138",
	expectedRating: 3,
}

func TestRating(t *testing.T) {
	// Arrange
	ratingMap := rating.GetRatingMap(&testCase.conf)
	order := &order.Order{
		UserID: testCase.userID,
	}
	kitchen := &kitchen.Kitchen{
		ID: testCase.kitchenID,
	}
	expectedRating := testCase.expectedRating

	// Act
	resultRating := ratingMap.GetOrderToKitchenRating(order, kitchen)

	// Assert
	assert.Equalf(t, expectedRating, resultRating, "Error! expected rating = %f, got %f", expectedRating, resultRating)
}
